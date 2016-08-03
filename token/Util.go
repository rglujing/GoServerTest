package token;

import (
	"encoding/json"
	l4g "code.google.com/p/log4go"
)

func Unmarshal_Request_Para(para string, v interface{}) (err * TokenErr) {
	
	tmp := []byte(para)
	length := len(tmp)
	/* para may be decrypt result, could contain '\x00' */
	for tmp[length-1] == 0 { 
		length--
	}
            
	e := json.Unmarshal([]byte(tmp[:length]), v)

	if e != nil {
		err = new(TokenErr)
		err.RetCode = RET_INVALID_JSON
		err.Msg = e.Error()
	}

	return
}


func  Marshal_Response_Para(req interface{}, err * TokenErr) (ret string) {
	
	var resp Common_ParaResp_Fail
	resp.Result = err.RetCode
	resp.Msg = err.Msg
	resp.ReqData = req

	msg, e := json.Marshal(resp)

	if e != nil {
		l4g.Error("Resp Fail json marshal failed %s", e.Error())
		return Internal_error_json
	} else {
		return string(msg)
	} 
}

