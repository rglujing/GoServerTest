package token;

import (
	"strings"
	"encoding/json"
	"strconv"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/etcd"
	"TokenServ/token/bdb"
)


type  refreshToken_ParaReq        struct {
	Appid     string          `json:"appid"`
	Secret    string          `json:"secret"`
	AsToken   string          `json:"AsToken"`
	AddExpire uint32          `json:"addExpire,omitempty"`
	ReqData   interface{}     `json:"reqdata,omitempty"`
}

type   refreshToken_ParaResp_Suc   struct {
	Result    int             `json:"result"`
	AsToken   string          `json:"AsToken"`
	Expireln  uint32          `json:"expireln"`
	Usertid   string          `json:"usertid"`
	ReqData   interface{}     `json:"req,omitempty"`
}


func RefreshToken  (para string, addr string) (ret  string) { 

	var reqpara refreshToken_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}
	
	if strings.EqualFold(reqpara.AsToken, "") {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}

	err, etnode := EntryControl(reqpara.Appid, addr, METHOD_REFRESH_TOKEN, true, reqpara.Secret)
	
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}


	ret, err = refresh_token( &reqpara , etnode)

	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	return
}

func refresh_token ( req * refreshToken_ParaReq, etnode * tokenetcd.Etcd_token_node ) ( ret string, err * TokenErr ) {
	var bdbReadWriter   BdbStore

	ret, e := tokenbdb.Bdb_server_get( BdbPrefix + req.AsToken )
	
	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_BDB_IO
		err.Msg = e.Error()
		return

	} 

	if strings.EqualFold(ret, "") {
		err = Err_Invalid_AsToken
		return
	}


	e = json.Unmarshal([]byte(ret), &bdbReadWriter)

	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return
	}

	if req.AddExpire <= 0 {
		req.AddExpire = etnode.DefaultExpire
	}

	bdbReadWriter.SetExpire = bdbReadWriter.SetExpire + req.AddExpire


	if bdbReadWriter.SetExpire > etnode.MaxExpire {
		err = new(TokenErr)
		err.RetCode = RET_EXIPIER_OVERFLOW
		err.Msg = Expire_Overflow + ",MAXEXIPRE=" + strconv.Itoa(int(etnode.MaxExpire))
		return
	}

	
	jstr, e := json.Marshal(bdbReadWriter)
	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return
	} else {
		e = tokenbdb.Bdb_server_set( BdbPrefix + req.AsToken, string(jstr) )	
	}


	if e != nil {
		l4g.Error("write bdb error %s", e.Error())
		err = new(TokenErr)
		err.RetCode = RET_INVALID_BDB_IO
		err.Msg = e.Error()
		return
	} else {
		var resp  refreshToken_ParaResp_Suc
		resp.Result = 0
		resp.AsToken =  req.AsToken
		resp.Expireln = bdbReadWriter.SetExpire
		resp.Usertid =  bdbReadWriter.Usertid
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
