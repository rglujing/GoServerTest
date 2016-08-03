package token;

import (
	"time"
	"crypto/md5"
	"crypto/rand"
	"strconv"
	"strings"
	"encoding/json"
	"encoding/hex"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/etcd"
	"TokenServ/token/redis"
	"TokenServ/token/bdb"
)

type requestSession_ParaReq  struct {
        Appid         string       `json:"appid"`
	Secret        string       `json:"secret"`
	AuCode        string       `json:"AuCode"`
	CallBackUrl   string       `json:"callbackurl"`
	CacheData     string       `json:"cachedata"`
	ReqData       interface{}  `json:"reqdata,omitempty"`
}

type requestSession_ParaResp_Suc struct {
	Result        int           `json:"result"`
	AsToken       string        `json:"AsToken"`
	Expireln      uint32        `json:"expireln"`
	Usertid       string        `json:"usertid"`
	ReqData       interface{}   `json:"req,omitempty"`
}

func RequestSession  (para string, addr string) (ret  string) { 

	var reqpara requestSession_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	if strings.EqualFold("", reqpara.AuCode) {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}


	err, etnode := EntryControl(reqpara.Appid, addr, METHOD_REQUEST_SESSION, true, reqpara.Secret)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}
	
	/*检查cache大小*/
	if len(reqpara.CacheData) > etnode.MaxCacheLen {
		err = Err_Cache_Overlen
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	ret, err = request_session ( &reqpara, etnode )

	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	return
	
}



func request_session ( req * requestSession_ParaReq, etnode * tokenetcd.Etcd_token_node) ( ret string, err * TokenErr ) {
	
	var redisReader RedisStore
	var bdbWriter   BdbStore

	ret, e := tokenredis.Redis_server_get( RedisPrefix + req.AuCode )
	
	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_REDIS_IO
		err.Msg = e.Error()
		return
	} 

	if strings.EqualFold(ret, "") {
		err = Err_Invalid_AuCode
		return
	}

	e = json.Unmarshal([]byte(ret), &redisReader)

	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return 
	}

	bdbWriter.Appid = redisReader.Appid
	bdbWriter.SetExpire = redisReader.SetExpire
	bdbWriter.Uname = redisReader.Uname
	bdbWriter.Usertid = redisReader.Usertid
	bdbWriter.CacheData = req.CacheData


	i := make([]byte, 4)
	_, e = rand.Read(i)
	if e != nil {
		l4g.Error("request_session rand error:%s", e.Error())
		err = Err_Internal
		return
	}
	
	tm := time.Now()
	createTM := tm.Unix()
	
	var j int
	j = int (i[0] )  + int (i[1]) + int (i[2] )
	j = j<<8
	j += int( i[3] )
	
	aucode := strconv.Itoa(j) + strconv.FormatInt(createTM, j%16+10)
	tmp :=  md5.Sum( []byte(aucode) ) 
	astoken := hex.EncodeToString ( tmp[:] )
	
	bdbWriter.CreateTM = createTM
	
	jstr, e := json.Marshal(bdbWriter)
	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return 
	
	} else {
		e = tokenbdb.Bdb_server_set( BdbPrefix + astoken, string(jstr) )	
	}


	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_BDB_IO
		err.Msg = e.Error()
		return 
	} else {
		/*Del Aucode*/
		tokenredis.Redis_server_del( RedisPrefix + req.AuCode )

		var resp  requestToken_ParaResp_Suc
		resp.Result = 0
		resp.AsToken = astoken
		resp.Expireln = bdbWriter.SetExpire
		resp.Usertid = bdbWriter.Usertid
		resp.ReqData = req.ReqData

		var msg []byte
		msg, e = json.Marshal(resp)
		if e != nil {
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
