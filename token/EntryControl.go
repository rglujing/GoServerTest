package token;

import (
	"time"
	"strings"
	"regexp"
	l4g "code.google.com/p/log4go"
	"TokenServ/token/etcd"
)

type VisitCountLimit struct {
	LastMin int64
	VisitCnt uint32
}

var (
	VisitCountControl = make(map[string]*VisitCountLimit)
)

func EntryControl(appid string, ip string, meth string, sec bool, secret string) (e * TokenErr,  etnode * tokenetcd.Etcd_token_node) {
	
	enable_ip     := false
	enable_meth   := false
	enable_secret := false

	if strings.EqualFold(appid, "") {
		etnode = tokenetcd.Etcd_Get_By_Ip(ip)
		if etnode != nil {
			enable_ip = true
		} else {
			e = Err_Invalid_Mix
			return 
		}
		
	} else {
		etnode = tokenetcd.Etcd_Get_By_Appid(appid)
		
		if etnode == nil {
			e = Err_Invalid_Appid
			return
		}


		allowips := etnode.AllowIPS
		for _, allowip := range allowips {
   
			b,e := regexp.MatchString(allowip, ip) 
			if e!=nil {
				l4g.Error("UnMatched IP %s , %s , %s", allowip, ip, e.Error())
				continue
			}

			if b { 
				l4g.Debug("Match IP %s, %s", allowip, ip)
				enable_ip = true
				break
			}
		}
		
	}
	

	if !enable_ip {
		e = Err_Invalid_GW
		return 
	}


	methods := etnode.OpenMethods

	for _,method := range methods {
		if strings.EqualFold(method, meth) {
			enable_meth = true
			break
		}
	}

	if !enable_meth {
		e = Err_Invalid_Method
		return 
	}
	

	if sec {
		if strings.EqualFold(secret, etnode.SvrPwd) {
			enable_secret = true
		} 
		if !enable_secret {
			e = Err_Invalid_Secret
			return 
		}
	}
	
	limit,ok := VisitCountControl[etnode.AppId]
	tm  := time.Now()
	min := tm.Unix()
	if ok {
		if  min - limit.LastMin < 60 {
			if limit.VisitCnt >= etnode.AccessLimit {
				e = Err_Visit_TooMany
				limit.VisitCnt ++
				return
			} else {
				limit.VisitCnt ++
			}
		} else {
			limit.LastMin = min
			limit.VisitCnt = 1
		}
	} else {
		limit = new(VisitCountLimit)
		limit.LastMin = min 
		limit.VisitCnt = 1
		VisitCountControl[etnode.AppId] = limit
	}

	e = nil
	return
}
