package tokenetcd;

import (
	/*
        "net"
        "net/url"
        "os"
	*/
        "fmt"
        "encoding/json"
        "testing"
	_"strconv"
	"time"
	l4g  "code.google.com/p/log4go"
)

func genJson(k int) string {

        var TokenList []Etcd_token_node
    	var t Etcd_token_node

        t = Etcd_token_node {
                AppId : "1.10001",
                Name : "test001",
                SvrPwd : "AAAAA",
                ReqEncode : "UTF-8",
                RespEncode : "UTF-8",
                AllowIPS: []string{ "127.1.[0-1].*:[1-9]+", "192.168.0.1", "127.0.0.1" },
                OpenMethods: []string{ "CreateAuCode", "RequestToken", "RequestSession", "RefreshToken", "SetSessionData" },
                AuCodeLife : 600,
                DefaultExpire : 3600,
                MaxExpire : 7200,
                AccessLimit : 10,
                MaxCacheLen : 50,
		MaxRequestLen: 200,
                AuthCallBackUrls : "http://www.baidu.com"}
        TokenList = append(TokenList, t)

        t = Etcd_token_node {
                AppId : "1.10002",
                Name : "test002",
                SvrPwd : "AAAAA",
                ReqEncode : "UTF-8",
                RespEncode : "UTF-8",
                AllowIPS: []string{ "192.168.0.1", "192.168.0.2" },
                OpenMethods: []string{ "CreateAuCode", "RequestToken" },
                AuCodeLife : 600,
                DefaultExpire : 3600,
                MaxExpire : 7200,
                AccessLimit : 10,
                MaxCacheLen : 512,
                AuthCallBackUrls : "http://www.baidu.com"}
        TokenList = append(TokenList, t)

        t = Etcd_token_node {
                AppId : "1.10003",
                Name : "test003",
                SvrPwd : "AAAAA",
                ReqEncode : "UTF-8",
                RespEncode : "UTF-8",
                AllowIPS: []string{ "192.168.0.1", "192.168.0.2" },
                OpenMethods: []string{ "CreateAuCode", "RequestToken" },
                AuCodeLife : 600,
                DefaultExpire : 3600,
                MaxExpire : 7200,
                AccessLimit : 10,
                MaxCacheLen : 512,
                AuthCallBackUrls : "http://www.baidu.com"}
        TokenList = append(TokenList, t)

        j,e := json.Marshal(TokenList)
        if e != nil {
                return ""
        }  

	return string(j)
}


func outputs(et * Etcd_token) {
	
	       for i , v := range et.data {
                fmt.Printf("%s", i)
                fmt.Printf("\n")
                fmt.Printf("%v", v)
                fmt.Printf("\n")
        }


}
/*
func TestEtcd(t *testing.T)  {

	l4g.LoadConfiguration("/opt/go/src/TokenServ/conf/testlog.xml")
	et,e := Etcd_init([]string{"https://10.15.201.55:4001", "https://10.15.201.55:4002"}, "ca/etcd.pem", "ca/etcd.key")

	if e != nil {
		t.Fatal(e)
	}
	var jsonstr string

	jsonstr = genJson(5)
	et.Etcd_Push("Testing", jsonstr)

	e = et.Etcd_Load("Testing")
	if e != nil {
		t.Fatal(e)
	}
	
	outputs(et)
	
	et.Etcd_Watch("/Testing")

	for {
	jsonstr = genJson(2)
	et.Etcd_Push("Testing", jsonstr)
	time.Sleep(1*time.Second)
	outputs(et)
	}
	
	time.Sleep(1*time.Second)
}
*/
func TestHelper(t *testing.T) {
	l4g.LoadConfiguration("/opt/go/src/TokenServ/conf/testlog.xml")
	et,e := Etcd_init([]string{"https://10.15.201.55:4001", "https://10.15.201.55:4002"}, "ca/etcd.pem", "ca/etcd.key")

	if e != nil {
		t.Fatal(e)
	}
	/*
	e = et.Etcd_Load("/tokenserv/tokendef")
	if e != nil {
		t.Fatal(e)
	}
	*/
	var jsonstr string

	jsonstr = genJson(5)
	et.Etcd_Push("/tokenserv/tokendef", jsonstr)

	e = et.Etcd_Load("/tokenserv/tokendef")
	if e != nil {
		t.Fatal(e)
	}
	
	outputs(et)
	
	fmt.Println("++++++++++++++++=")
	tmp := et.Etcd_Get_By_Appid("1.10003")
	fmt.Println(tmp.AppId)

	time.Sleep(5*time.Second)

}
