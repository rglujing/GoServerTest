package token;

import (
	"strings"
	"encoding/json"
	l4g "code.google.com/p/log4go"
	_"TokenServ/token/etcd"
	"TokenServ/token/bdb"
)

type   setSessionData_ParaReq  struct {
        Appid         string       `json:"appid"`
	Secret        string       `json:"secret"`
	AsToken       string       `json:"AsToken"`
	CacheData     string       `json:"cachedata"`
	ReqData       interface{}  `json:"reqdata,omitempty"`
}

type   setSessionData_ParaResp_Suc struct {
	Result        int           `json:"result"`
	AsToken       string        `json:"AsToken"`
	Expireln      uint32        `json:"expireln"`
	Usertid       string        `json:"usertid"`
	ReqData       interface{}   `json:"req,omitempty"`
}

func  SetSessionData (para string, addr string) (ret  string) {

	var reqpara setSessionData_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)

	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}
	
	if strings.EqualFold("", reqpara.AsToken) {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}

	err, etnode := EntryControl(reqpara.Appid, addr, METHOD_SET_SESSION_DATA, true, reqpara.Secret)	
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	/*检查包大小*/
	if len(reqpara.CacheData) > etnode.MaxCacheLen {
		err = Err_Cache_Overlen
		return Marshal_Response_Para(reqpara.ReqData, err)
	}


	ret, err = set_session_data( &reqpara )
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	return	

}

func set_session_data ( req * setSessionData_ParaReq ) ( ret string, err * TokenErr ) {
	
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

	bdbReadWriter.CacheData = req.CacheData
	
	jstr, e := json.Marshal(bdbReadWriter)
	if e != nil {
		l4g.Error("Json marshal error %s", e.Error())
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return
	
	} else {
		e = tokenbdb.Bdb_server_set( BdbPrefix + req.AsToken, string(jstr) )	
	}


	if e != nil {
		l4g.Error("write Bdb error %s", e.Error())
		err = new(TokenErr)
		err.RetCode = RET_INVALID_BDB_IO
		err.Msg = e.Error()
		return
	} else {
		var resp  setSessionData_ParaResp_Suc
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
