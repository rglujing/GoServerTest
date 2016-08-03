package TcpRoute;

import (
	"testing"
	l4g  "code.google.com/p/log4go"
	)

func TestListen(t * testing.T) {
	
	l4g.LoadConfiguration("/opt/go/src/TokenServ/conf/testlog.xml")

	e := Token_tcp_listen(":8188")
	
	if e != nil {
		l4g.Error(e.Error())	
	}

}
