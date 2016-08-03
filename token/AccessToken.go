package token;

import (
	"time"
	"strings"
	"encoding/json"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/bdb"
)

type  accessToken_ParaReq        struct {
	AsToken    string          `json:"AsToken"`
	ReqData    interface{}     `json:"req,omitempty"`
}

type  accessToken_ParaResp_Suc  struct {
	Result     int             `json:"result"`
	Uname      string          `json:"uname"`
	Usertid    string          `json:"usertid"`
	Expireln   uint32          `json:"expireln"`
	CacheData  interface{}     `json:"cachedata,omitempty"`
	ReqData    interface{}     `json:"req,omitempty"`
}

func AccessToken (para string, addr string) (ret  string) {
	
	var reqpara accessToken_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}	
	
	if strings.EqualFold("", reqpara.AsToken) {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}


	ret, err = access_token( &reqpara )
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	} 

	return 
}


func access_token( req * accessToken_ParaReq ) ( ret string, err * TokenErr ) {
	var bdbReader   BdbStore

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

	e = json.Unmarshal([]byte(ret), &bdbReader)

	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
		return
	}

	tm := time.Now()
	nowTM := tm.Unix()
	
	if nowTM >  ( bdbReader.CreateTM + int64(bdbReader.SetExpire) ) {
		err = Err_AsToken_Timeout
		return
	} else {
		var resp accessToken_ParaResp_Suc
		resp.Result = 0
		resp.Uname = bdbReader.Uname
		resp.Usertid = bdbReader.Usertid
		resp.Expireln = bdbReader.SetExpire
		resp.CacheData = bdbReader.CacheData

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


