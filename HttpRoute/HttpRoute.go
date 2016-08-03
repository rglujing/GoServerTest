package HttpRoute;

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	l4g "code.google.com/p/log4go"
	"TokenServ/token"
	"TokenServ/crypt"
	)


var handlerMap map[string]int

func Token_http_reg(port string, pat map[string]string, cer string, key string) {
	var err error

	l4g.Debug("HTTP Server Listen at %s Begin", port)

	handlerMap = make(map[string]int)

	http.HandleFunc("/", http_handler)

	handlerMap[ pat["CreateAuCode"] ] = token.CREATE_AUCODE
	handlerMap[ pat["RequestToken"] ] = token.REQUEST_TOKEN
	handlerMap[ pat["AccessToken"]  ] = token.ACCESS_TOKEN
	handlerMap[ pat["RefreshToken"] ] = token.REFRESH_TOKEN
	handlerMap[ pat["InspectToken"] ] = token.INSPECT_TOKEN
	handlerMap[ pat["DestroyToken"] ] = token.DESTROY_TOKEN
	handlerMap[ pat["RequestSession"] ] = token.REQUEST_SESSION
	handlerMap[ pat["SetSessionData"] ] = token.SET_SESSION_DATA

	if strings.EqualFold("", cer) || strings.EqualFold("", key) {
		l4g.Error("no valid SLL CA FILES")
		return
	} else {
		err = http.ListenAndServeTLS(port, cer, key, nil)
	}

	if err != nil {
		l4g.Error("Init http failed, %s", err.Error())
		return
	}
}


func http_parse_request( r * http.Request) (res string, e error) {
	/*Post is prefer*/
	if r.ContentLength > 0 {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			e = err
			return 
		}
	
		v := r.Header.Get("Content-Type")
		/*DeCrypt*/
		if strings.EqualFold(v, "application/encrypt") {
			res = tokenaes.Decrypt(b)
		} else {
			res = string(b)
		}
			

	} else {
		r.ParseForm()
		for i,_ := range r.Form {
			res = i
			break
		}
		l4g.Debug("GET para is", res)
	}

	return
}

func http_handler(w http.ResponseWriter, r * http.Request) {
	
	l4g.Debug("Http request method is %s\n", r.URL.Path)

	if (token.MAX_JSON_LEN != 0) && (r.ContentLength > (int64)(token.MAX_JSON_LEN) )  {
		fmt.Fprintf(w, token.Conten_tlong_json)
		return
	}

	str, e := http_parse_request(r)

	if e != nil {
		l4g.Error("Parse parameter %s", e.Error())
		fmt.Fprintf(w, token.Invalid_method_json)
	} else {
		index, exist := handlerMap[r.URL.Path]
		if exist {
			ret := token.Entry_Filter(str, r.RemoteAddr, index)
			fmt.Fprintf(w, ret)
		} else {
			fmt.Fprintf(w, token.Invalid_method_json)
		}
	}


}
