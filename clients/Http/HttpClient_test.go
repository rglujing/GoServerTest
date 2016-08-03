package HttpClient;


import (
	"fmt"
	"testing"
	"encoding/json"
	"time"
	)

func TestAuCode( t * testing.T ) {
	
        var obj map[string]string
        obj = make(map[string]string)
        obj["f"]="XXX"
        obj["q"]="QQQ"
        para := CreateAuCode_ParaReq{Appid:"1.10001", Callbackurl:"www.qq.com", SetExpire:10, Uname:"XXX", Usertid:"001", ReqData:obj}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	CreateAuCode(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())

}



func TestReqToken1( t * testing.T ) {
	
	para := RequestToken_ParaReq{Appid:"1.10001", Secret:"AAAAA", AuCode:"XX951071962571a1"}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	RequestToken(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())
}


func TestReqToken2( t * testing.T ) {
	
	para := RequestToken_ParaReq{Appid:"1.10001", Secret:"AAAAA", AuCode:"149196c9gdfeg"}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	RequestToken(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())
}



func TestReqSession( t * testing.T ) {
	

	para := RequestSession_ParaReq{Appid:"1.10001", Secret:"AAAAA", AuCode:"769915jbiajl", CacheData:"XXXXXXXX"}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	RequestSession(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())
}


/*
func TestSetSession( t * testing.T ) {
	

	para := SetSessionData_ParaReq{Appid:"1.10001", Secret:"AAAAA", AsToken:"d3ee4ab3efeea2095c53f5d962ec4161", CacheData:"VVVVVVVVVVVVVVV"}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	SetSessionData(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())
}



func TestRefreshToken( t * testing.T ) {
	
        var obj map[string]int
        obj = make(map[string]int)
        obj["f"]=9
        obj["q"]=1
 

	para := RefreshToken_ParaReq{Appid:"1.10001", Secret:"AAAAA", AsToken:"d3ee4ab3efeea2095c53f5d962ec4161", AddExpire:500}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	RefreshToken(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())
}




func TestAccessToken( t * testing.T ) {
	
	para :=AccessToken_ParaReq{AsToken:"d3ee4ab3efeea2095c53f5d962ec4161"}
	
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	AccessToken(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())

}


func TestInspectToken( t * testing.T ) {
	
	para := InspectToken_ParaReq{AsToken:"d3ee4ab3efeea2095c53f5d962ec4161"}
	
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	InspectToken(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())

}



func TestDestroyToken( t * testing.T ) {
	
	para := InspectToken_ParaReq{AsToken:"d3ee4ab3efeea2095c53f5d962ec4161"}
        msg,err := json.Marshal(para)
        if err != nil {
                fmt.Printf("err is %s\n", err.Error())
                return
        }   
        msgstr :=  string(msg)
        fmt.Printf("MSGS is %s\n", msgstr)
	tm := time.Now()
	fmt.Println(tm.Unix())
	DestroyToken(false, true,  msgstr)
	tm = time.Now()
	fmt.Println(tm.Unix())

}
*/
