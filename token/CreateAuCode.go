package token;

import (
	"time"
	_"crypto/md5"
	"crypto/rand"
	"strconv"
	"strings"
	"encoding/json"
	_"encoding/hex"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/etcd"
	"TokenServ/token/redis"
	_"TokenServ/token/bdb"
)

type createAuCode_ParaReq struct {
        Appid         string       `json:"appid"`
        Callbackurl   string       `json:"callbackurl"`
        SetExpire     uint32       `json:"setExpire,omitempty"`
        Uname         string       `json:"uname"`
        Usertid       string       `json:"usertid"`
        ReqData       interface{}  `json:"reqdata,omitempty"`
}

type CreateAuCode_ParaResp_Suc struct {
	Result   int                 `json:"result"`
	AuCode   string              `json:"AuCode"`
	ReqData  interface{}         `json:"req,omitempty"`
}


func CreateAucode (para string, addr string) (ret  string) {

	var reqpara createAuCode_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	if strings.EqualFold("", reqpara.Uname) || strings.EqualFold("", reqpara.Usertid) {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}


	/*Access Control*/
	err, etnode := EntryControl(reqpara.Appid, addr, METHOD_CREATE_AUCODE, false, "")
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}
	
	ret, err = create_aucode( &reqpara, etnode )
	/*Interl error, do not show to invoker*/
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}
	
	return ret
}

func create_aucode( req * createAuCode_ParaReq, etnode * tokenetcd.Etcd_token_node) (ret string, err *TokenErr) {
	
	var redisWriter RedisStore

	redisWriter.Appid = etnode.AppId

	if req.SetExpire == 0 {
		redisWriter.SetExpire = etnode.DefaultExpire
	} else {
		redisWriter.SetExpire = req.SetExpire
	}

	redisWriter.Uname = req.Uname
	redisWriter.Usertid = req.Usertid

	i := make([]byte, 4)
	_, e := rand.Read(i)
	if e != nil {
		l4g.Error("create_aucode rand error:%s", e.Error())
		err = Err_Internal
		return
	}
	
	tm := time.Now()
	createTM := tm.Unix()
	redisWriter.CreateTM = createTM
	
	var j int
	j = int (i[0] )  + int (i[1]) + int (i[2] )
	j = j<<8
	j += int( i[3] )
	
	aucode := strconv.Itoa(j) + strconv.FormatInt(createTM, j%16+10)
	
	jstr, e := json.Marshal(redisWriter)
	if e != nil {
		err = new(TokenErr)	
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return
	} else {
		e = tokenredis.Redis_server_set( RedisPrefix + aucode, string(jstr), etnode.AuCodeLife)
	}


	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_REDIS_IO
		err.Msg = e.Error()
		return
	} else {
		var resp CreateAuCode_ParaResp_Suc
		resp.Result = RET_SUC
		resp.AuCode =  aucode
		resp.ReqData = req.ReqData
		var msg []byte
		msg, e = json.Marshal(resp)
		if e != nil {
			l4g.Error("Resp Suc json marshal failed %s", e.Error())
			err = new(TokenErr)
			err.RetCode = RET_INVALID_JSON
			err.Msg = e.Error()
			return
		} else {
			err = nil
			ret = string(msg)
			return
		}
	}
}
