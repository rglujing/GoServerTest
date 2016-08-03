package HttpRoute;

import (
	"testing"
	l4g  "code.google.com/p/log4go"
	"time"
	)

func TestListen(t * testing.T) {
	
	l4g.LoadConfiguration("/opt/go/src/TokenServ/conf/testlog.xml")

	var pat map[string]string

	pat = make(map[string]string)
	pat["CreateAuCode"] = "/CreateAuCode"
	pat["RequestToken"] = "/RequestToken"
	pat["AccessToken"]  = "/AccessToken"
	pat["RefreshToken"] = "/RefreshToken"
	pat["InspectToken"] = "/InspectToken"
	pat["DestroyToken"] = "/DestroyToken"
	pat["RequestSession"] = "/RequestSession"

	//Token_http_reg(":8188", pat, "./server.cer", "./server.key")
	l4g.Debug("Before")
	Token_http_reg(":8188", pat, "/", "/")
	l4g.Debug("Next")

	for {
		time.Sleep(time.Second)
	}
}
