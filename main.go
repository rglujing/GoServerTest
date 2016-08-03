package main;

import (
	"os"
	"fmt"
	"time"
	"encoding/json"
	"TokenServ/HttpRoute"
	"TokenServ/TcpRoute"
	"TokenServ/token"
	l4g "code.google.com/p/log4go"
	)

type Config struct {
	HttpPort      string
	UrlPattern    map[string]string
	HttpServercrt string
	HttpServerkey string
	TcpPort       string
	TokenConf     string
}

func loadconf() (ret *Config) {

	Conf := new(Config)

	file,_ := os.Open("./conf/Interface.json")
	decoder := json.NewDecoder(file)
	Conf.UrlPattern = make(map[string]string)
	err:=decoder.Decode(Conf)
	if err != nil {
		l4g.Error("load config failed %s", err.Error())
	}

	l4g.Info("HttpPort is %s",      Conf.HttpPort      )
	l4g.Info("UrlPattern is %v",    Conf.UrlPattern    )
	l4g.Info("HttpServercrt is %s", Conf.HttpServercrt )
	l4g.Info("HttpServerkey is %s", Conf.HttpServerkey )
	l4g.Info("TcpPort is %s",       Conf.TcpPort       )
	l4g.Info("TokenConf is %s",     Conf.TokenConf     )	

	for _,v := range Conf.UrlPattern {
		l4g.Info(v)
	}

	return Conf
}


func main() {
	
	l4g.LoadConfiguration("./conf/log.xml")

	Conf := loadconf()
	
	e := token.Token_Init(Conf.TokenConf)
	
	if e != nil {
		fmt.Printf("Token init failed %s\n", e.Error())
		return
	}


	go HttpRoute.Token_http_reg(Conf.HttpPort, Conf.UrlPattern, Conf.HttpServercrt, Conf.HttpServerkey)
	go TcpRoute.Token_tcp_listen(Conf.TcpPort)

	for {
		time.Sleep(time.Second)	
	}
}
