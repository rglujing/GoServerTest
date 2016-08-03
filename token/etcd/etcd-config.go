package tokenetcd;

import (
	"github.com/coreos/go-etcd/etcd"
	"encoding/json"
	l4g "code.google.com/p/log4go"
	"regexp"
	"fmt"
	"strings"
	"sync"
	)

type Etcd_token_node struct {
	AppId             string    `json:"appid"`
	Name              string    `json:"name"`
	SvrPwd            string    `json:"svrpwd"`
	ReqEncode         string    `json:"reqencode"`
	RespEncode        string    `json:"respencode"`
	AllowIPS          []string	`json:"allowips"`
	OpenMethods       []string	`json:"openmethod"`
	AuCodeLife        int		`json:"AuCodeLife"`
	DefaultExpire     uint32	`json:"defaultExpire"`
	MaxExpire         uint32	`json:"maxExpire"`
	AccessLimit       uint32	`json:"accessLimit"`
	MaxCacheLen       int		`json:"maxCacheLen"`
	MaxRequestLen     int		`json:"maxRequestLen"`
	AuthCallBackUrls  string	`json:"AuthCallBackUrls"`
}

type Etcd_token_map map[string]Etcd_token_node


type Etcd_token struct {
	Mark string
	cli   *etcd.Client
	index uint64         // for watch function, since index
	data  Etcd_token_map // map[appid] tp Etcd_token_node
	MaxRequestLen  int
	NodeName string
	mutex1   sync.Mutex
	mutex2   sync.Mutex
}

var defaultet * Etcd_token;

func Etcd_init(machines []string, cert, key string) ( et * Etcd_token, e error){
	
	et = new(Etcd_token)
	
	if  strings.EqualFold(cert, "") {
		et.cli = etcd.NewClient(machines)
	} else {
		et.cli,e = etcd.NewTLSClient(machines, cert, key, "")
	}
	
	if e != nil {
		et = nil
		defaultet = nil
		return
	}
	
	et.data = make(Etcd_token_map)
	et.Mark = "?????????"
	defaultet = et
	return	
}

func Etcd_get_max_request_len() (int) {
	return defaultet.MaxRequestLen
}

func Etcd_Load(key string) (e error) {
	return defaultet.Etcd_Load(key)
}

func Etcd_Push(key string, val string) (e error) {
	return defaultet.Etcd_Push(key, val)
}

func Etcd_Get_By_Appid(appid string) (node * Etcd_token_node) {
	return defaultet.Etcd_Get_By_Appid(appid)
}

func Etcd_Get_By_Ip(ip string) (node * Etcd_token_node) {
	return defaultet.Etcd_Get_By_Ip(ip)
}

func Etcd_Watch(key string) (e error) {
	return defaultet.Etcd_Watch(key)
}


func (et * Etcd_token) Etcd_Load(key string) (e error) {
	
	r, e := et.cli.Get(key, false, false)
	if e != nil {
		return
	}
	et.index = r.Node.ModifiedIndex+1
	e = et.etcd_refresh_data( []byte(r.Node.Value) )
	return
}


func (et * Etcd_token) etcd_refresh_data (buf []byte)  (e error) {
	
	et.mutex1.Lock()
	et.mutex2.Lock()

	defer et.mutex1.Unlock()
	defer et.mutex2.Unlock()


	for i,_ := range et.data {
		delete(et.data, i)
	}


	var t []Etcd_token_node
	e = json.Unmarshal( buf , &t )

	if e != nil {
		l4g.Debug("Analyst data failed!!! %s", e)
		return
	}

	for _ , v := range t {
		et.data[v.AppId] = v
		if v.MaxRequestLen > et.MaxRequestLen {
			et.MaxRequestLen = v.MaxRequestLen
		}
	}
	
	return
}

func (et * Etcd_token) Etcd_Push(key string, val string) (e error) {
	
	_, e = et.cli.Set(key, val, 0)
	return 
}

func (et * Etcd_token) Etcd_Get_By_Appid(appid string) (node * Etcd_token_node) {
	
	et.mutex2.Lock()
	defer et.mutex2.Unlock()
	tnode , ok := et.data[appid]

	if ok  {
		node = &tnode
		return
	} else {
		node = nil
		return
	}	

}

func (et * Etcd_token) Etcd_Get_By_Ip(ip string) (node * Etcd_token_node) {

	et.mutex1.Lock()
	defer et.mutex1.Unlock()
	node = nil	

	for _, tnode := range et.data {
		
		allowips := tnode.AllowIPS
		for _, allowip := range allowips {
			
			b,e := regexp.MatchString(allowip, ip)
			if e!=nil {
				l4g.Error("UnMatch IP %s %s %s", allowip, ip, e.Error())
				continue
			} 

			if b {
				l4g.Debug("Match IP %s %s", allowip, ip)
				node = &tnode
				return
			}
			
		}

	}

	return
	
}

func (et * Etcd_token) Etcd_Watch(key string) (e error) {
	
	receiver := make(chan *etcd.Response)
	stop := make(chan bool)
	
	et.NodeName = key
	
	go etcd_listener(et, receiver, stop)
	go et.cli.Watch(key, et.index, false, receiver, stop)

	return
}

func etcd_listener( et * Etcd_token, rec chan * etcd.Response, stop chan bool) {
	
	for {
		fmt.Printf("Listerner is working\n")
		t := <-rec
		
		if t == nil {
			/*May this close trigger panic?*/
			close(stop)
			fmt.Printf("Listerner is stopping\n")
			break
		}
		
		if t.Node != nil {
			l4g.Debug("t.Node.ModifiedIndex is %d", t.Node.ModifiedIndex)
			l4g.Debug("Update : t.Node.Value")
			l4g.Debug(t.Node.Value)
			et.index = t.Node.ModifiedIndex + 1
			e := et.etcd_refresh_data( []byte(t.Node.Value) )
			if e != nil {
				l4g.Debug("listener failed, can not refresh data")
			}
		} else {
			l4g.Debug("t.Node is nil")
		}
		
	}

	et.Etcd_Watch(et.NodeName)
}

