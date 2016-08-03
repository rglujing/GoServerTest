package token;

import (
	"os"
	"encoding/json"
	_"strings"
	_"regexp"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/etcd"
	"TokenServ/token/redis"
	"TokenServ/token/bdb"
	)

type Config struct {
	BdbServers    []string
	RedisServers  []string
	EtcdServers   []string
	EtcdNode      string
	EtcdCrt       string
	EtcdKey       string
}

func loadconf(filename string) (ret * Config){

	Conf := new(Config)
	file,_ := os.Open(filename)
	decoder := json.NewDecoder(file)
	err:=decoder.Decode(Conf)
	if err != nil {
		l4g.Error("load config failed %s", err.Error())
	}

	l4g.Info("BdbServers is %v",    Conf.BdbServers    )
	l4g.Info("RedisServers is %v",  Conf.RedisServers  )
	l4g.Info("EtcdServers is %v",   Conf.EtcdServers   )
	l4g.Info("EtcdNode is %v",      Conf.EtcdNode      )
	l4g.Info("EtcdCrt     is %s",   Conf.EtcdCrt       )
	l4g.Info("EtcdKey     is %s",   Conf.EtcdKey       )

	return Conf
}


func Token_Init( filename string ) (e error){

	Conf := loadconf(filename)

	e = token_etcd_init(Conf.EtcdServers, Conf.EtcdNode, Conf.EtcdCrt, Conf.EtcdKey)
	
	if e != nil {
		l4g.Error("Init etcd error: %s", e.Error())
		return
	}	

	_, e = tokenredis.Redis_init(Conf.RedisServers)
	if e != nil {
		l4g.Error("Init Redis error: %s", e.Error())
		return
	}	

	_, e = tokenbdb.Bdb_init(Conf.BdbServers)
	if e != nil {
		l4g.Error("Init Bdb error: %s", e.Error())
		return
	}	

	return
}


func token_etcd_init( servers []string, nodename, cert, key string ) ( e error ) {

	_, e  = tokenetcd.Etcd_init( servers, cert, key )
	
	if e != nil {
		return
	}	
	
	e = tokenetcd.Etcd_Load(nodename)
	if e != nil {
		return
	}

	e = tokenetcd.Etcd_Watch(nodename)
	if e != nil {
		return
	}

	MAX_JSON_LEN = tokenetcd.Etcd_get_max_request_len()
	return
}


func Entry_Filter(para string, addr string, index int) (ret  string) {
	switch index {
		case  CREATE_AUCODE:
			ret = CreateAucode(para , addr)		
		case  REQUEST_TOKEN: 
			ret = RequestToken(para , addr)
		case  ACCESS_TOKEN:
			ret = AccessToken(para, addr)	
		case  REFRESH_TOKEN:
			ret = RefreshToken(para, addr)		
		case  INSPECT_TOKEN:
			ret = InspectToken(para, addr)
		case  DESTROY_TOKEN:
			ret = DestroyToken(para, addr)
		case  REQUEST_SESSION:
			ret = RequestSession(para, addr)
		case  SET_SESSION_DATA:
			ret = SetSessionData(para, addr)
		default:
			ret = Invalid_method_json
	}
	
	return 	

}
