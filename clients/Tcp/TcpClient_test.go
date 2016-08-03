package TcpClient;


import (
	"fmt"
	"testing"
	"encoding/json"
	"time"
	)

type Parameter struct {
        Appid         string     `json:"appid"`
        Callbackurl   string     `json:"callbackurl"`
        SetExpire     int        `json:"setExpire,omitempty"`
        Uname         string     `json:"uname"`
        Usertid       string     `json:"usertid"`
        Reqdata       interface{}  `json:"reqdata"`
}

func TestTcp( t * testing.T ) {
	
        var obj map[string]string
        obj = make(map[string]string)
        obj["f"]="XXX"
        obj["q"]="QQQ"
        para := Parameter{Appid:"1.10001", Callbackurl:"www.qq.com", SetExpire:10, Uname:"XXX你好", Usertid:"001", Reqdata:obj}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	CreateAuCode(msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())
	
}

