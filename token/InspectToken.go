package token;

import (
	"time"
	"strings"
	"encoding/json"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/bdb"
)


type  inspectToken_ParaReq        struct {
	AsToken    string          `json:"AsToken"`
	ReqData    interface{}     `json:"req,omitempty"`
}

type   inspectToken_ParaResp_Suc   struct {
	Result    int             `json:"result"`
	Appid     string          `json:"appid"`
	Usertid   string          `json:"usertid"`
	CreateTM  int64           `json:"createTM"`
	Expireln  uint32          `json:"expireln"`
	ReqData   interface{}     `json:"req,omitempty"`
}

func InspectToken (para string, addr string) (ret  string) {
	var reqpara inspectToken_ParaReq

	err := Unmarshal_Request_Para(para, &reqpara)
	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	}

	if strings.EqualFold("", reqpara.AsToken) {
		return Marshal_Response_Para(reqpara.ReqData, Err_Invalid_Param)
	}

	ret, err = inspect_token( &reqpara )

	if err != nil {
		return Marshal_Response_Para(reqpara.ReqData, err)
	} 
	
	return
}


func inspect_token  ( req * inspectToken_ParaReq ) ( ret string, err * TokenErr ) {
	var bdbReader   BdbStore

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
		var resp inspectToken_ParaResp_Suc
		resp.Result = 0
		resp.Appid = bdbReader.Appid
		resp.Usertid = bdbReader.Usertid
		resp.CreateTM = bdbReader.CreateTM
		resp.Expireln = bdbReader.SetExpire
	
		
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


