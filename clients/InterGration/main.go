package main

import (
	"os"
	"fmt"
	"flag"
	"encoding/json"
	"strings"
	"TokenServ/clients/InterGration/refer"
)
const (
	MTIP =  "set method,  it can be CreateAucode/C/CA    \n" + 
				"\t\t   RequestSession/RS    \n" + 
				"\t\t   RequestToken/RT      \n" + 
				"\t\t   RefreshToken/RF      \n" +
	 			"\t\t   AccessToken/A        \n" + 
				"\t\t   InspectToken/I       \n" +
				"\t\t   DestroyToken/D       \n" + 
				"\t\t   SetSessionData/S/SD  \n" +
				"\t\t   for example: \"-m SD\" , set method to be SetSessionData  \n"
)

type Config struct {
	Method    string
	Proto     string  `json:"proto"`
	ReqData   string  `json:"reqdata"`
	CacheData string  `json:"cache"`
	Appid     string  `json:"appid"`
	Uname     string  `json:"uname"`
	Usertid   string  `json:"usertid"`
	Secret    string  `json:"secret"`
	AuCode    string  
	AsToken   string 
	Expire    int
	Http      string  `json:"httpprefix"`
	Tcp       string  `json:"tcp"`
}

var conf *Config

func loadconf(filename string) () {
	
	conf = new(Config)

	f,err := os.Open( filename )
	if err != nil {
		fmt.Printf("err Open conf file %s failed\n", filename)
		os.Exit(-1)
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(conf)

	if err != nil {
		fmt.Printf("err decode conf file failed\n")
		os.Exit(-1)
	}

	return
}

func main() {
	
	method    := flag.String("m", "", MTIP )
	proto     := flag.String("proto",   "X", "use http or tcp" )
	reqdata   := flag.String("reqdata", "X",     "reqdata"         )
	cache     := flag.String("cache",   "X",     "cache data"      )	
	appid     := flag.String("appid",   "X",     "appid"      )	
	aucode    := flag.String("aucode",  "",  "aucode")
	astoken   := flag.String("astoken",  "", "astoken")
	expire    := flag.Int("expire", 0, "expire time")	

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:%s [OPTIONS], method must be assign \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	loadconf("./conf.json")
	
	if strings.EqualFold("", *method) {
		flag.Usage()
		os.Exit(-1)
	}
	
	if !strings.EqualFold("X", *proto) {
		conf.Proto = *proto
	}

	if !strings.EqualFold("X", *reqdata) {
		conf.ReqData = *reqdata
	}
	
	if !strings.EqualFold("X", *cache) {
		conf.CacheData = *cache
	}

	if !strings.EqualFold("X", *appid) {
		conf.Appid = *appid
	}

	conf.Expire = *expire
	conf.AuCode = *aucode
	conf.AsToken = *astoken

	switch *method {
		case "CreateAuCode", "C", "CA":
			conf.Method = "CreateAuCode"
			CreateAuCode()
		case "RequestToken", "RT":
			conf.Method = "RequestToken"
			RequestToken()
		case "RequestSession", "RS":
			conf.Method = "RequestSession"
			RequestSession()
		case "AccessToken", "A":
			conf.Method = "AccessToken"
			AccessToken()
		case "InspectToken", "I":
			conf.Method = "InspectToken"
			InspectToken()
		case "DestroyToken", "D":
			conf.Method = "DestroyToken"
			DestroyToken()
		case "RefreshToken", "RF":
			conf.Method = "RefreshToken"
			RefreshToken()
		case "SetSessionData", "S", "SD":
			conf.Method = "SetSessionData"
			SetSessionData()
		default:
			flag.Usage()
	}
}



func CreateAuCode() {
	
	var para refer.CreateAuCode_ParaReq
	para.Appid = conf.Appid
	para.SetExpire = uint32(conf.Expire)
	para.Uname = conf.Uname
	para.Usertid = conf.Usertid
	para.ReqData = conf.ReqData
	
	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_CREATE_AUCODE_REQ)
	}

}


func RequestToken() {
	var para refer.RequestToken_ParaReq
	para.Appid = conf.Appid
	para.Secret = conf.Secret
	para.AuCode = conf.AuCode
	para.ReqData = conf.ReqData
	
	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_REQUEST_TOKEN_REQ)
	}
}

func RefreshToken() {
	var para refer.RefreshToken_ParaReq

	para.Appid = conf.Appid
	para.Secret = conf.Secret
	para.AsToken = conf.AsToken
	para.AddExpire = uint32(conf.Expire)
	para.ReqData = conf.ReqData

	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_REFRESH_TOKEN_REQ)
	}

}

func AccessToken() {
	var para refer.AccessToken_ParaReq
	
	para.AsToken = conf.AsToken
	para.ReqData = conf.ReqData

	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_ACCESS_TOKEN_REQ)
	}

	
}

func InspectToken() {
	var para refer.InspectToken_ParaReq
	
	para.AsToken = conf.AsToken
	para.ReqData = conf.ReqData

	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_INSPECT_TOKEN_REQ)
	}

	
}

func DestroyToken() {
	var para refer.DestroyToken_ParaReq
	
	para.AsToken = conf.AsToken
	para.ReqData = conf.ReqData

	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_DESTROY_TOKEN_REQ)
	}

	
}

func RequestSession() {
	var para refer.RequestSession_ParaReq
	para.Appid = conf.Appid
	para.Secret = conf.Secret
	para.AuCode = conf.AuCode
	para.CacheData = conf.CacheData
	para.ReqData = conf.ReqData

	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_REQUEST_SESSION_REQ)
	}
}

func SetSessionData() {
	var para refer.SetSessionData_ParaReq
	para.Appid = conf.Appid
	para.Secret = conf.Secret
	para.AsToken = conf.AsToken
	para.CacheData = conf.CacheData
	para.ReqData = conf.ReqData

	ret, e := json.Marshal(&para)

	if e != nil {
		fmt.Printf("json marshal error : %s\n", e.Error())
	}

	switch conf.Proto {
		case "http":
			httpRequest(string(ret), conf.Method)
		case "tcp":
			tcpRequest(string(ret), AU_SET_SESSION_DATA_REQ)
	}

	
}
