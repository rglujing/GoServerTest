package token;

import (
	"strings"
	"encoding/json"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/bdb"
)


type destroyToken_ParaReq     struct {
	AsToken    string          `json:"AsToken"`
	ReqData    interface{}     `json:"req,omitempty"`
}

type destroyToken_ParaResp_Suc  struct {
	Result    int            `json:"result"`
	ReqData   interface{}    `json:"req,omitempty"`
}



func  DestroyToken (para string, addr string) (ret  string) {
	var reqpara destroyToken_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	if strings.EqualFold("", reqpara.AsToken) {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}

	ret,err = destroy_token ( &reqpara )
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	} 

	return
}


func destroy_token ( req * destroyToken_ParaReq ) ( ret string, err * TokenErr ) {

	ret, e := tokenbdb.Bdb_server_get( BdbPrefix + req.AsToken )
	
	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_BDB_IO
		err.Msg = e.Error()
		return
	} 

	if  strings.EqualFold(ret, "") {
		err = Err_Invalid_AsToken
		return
	}


	e = tokenbdb.Bdb_server_del( BdbPrefix + req.AsToken )
	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_BDB_IO
		err.Msg = e.Error()
		return
	} 

	var resp destroyToken_ParaResp_Suc
	resp.Result = 0
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


