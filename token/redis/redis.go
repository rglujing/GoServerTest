package tokenredis;

import (
	"github.com/keimoon/gore"
	"errors"
	"strings"
	_"fmt"
	)

type RedisToken struct {
	server_list []string
	server_using string
	conn * gore.Conn
}

var (
	NOVALIDSERVER = errors.New("invalid redis server list")
)

var defaultServer * RedisToken

func  Redis_server_set(key string, val string, tm int) (e error) {

	e = defaultServer.Redis_server_set(key,  val, tm)
	if e != nil {
		defaultServer.Redis_switch_server()
	} else {
		return
	}

	if defaultServer.conn != nil {
		return defaultServer.Redis_server_set(key,  val, tm)
	} else {
		return
	}
}

func  Redis_server_get(key string) ( ret string, e error) {

	ret, e = defaultServer.Redis_server_get(key)

	if e != nil {
		defaultServer.Redis_switch_server()
	} else {
		return
	}

	if defaultServer.conn != nil {
		return defaultServer.Redis_server_get(key)
	} else {
		return
	}
}

func Redis_server_del(key string) (ret string, e error) {

	ret, e = defaultServer.Redis_server_del(key)

	if e != nil {
		defaultServer.Redis_switch_server()
	} else {
		return
	}
	
	if defaultServer.conn != nil { 
		return defaultServer.Redis_server_del(key)
	} else {
		return
	}
	
}

func Redis_init( list []string ) (server * RedisToken, e error) {
	
	server = new(RedisToken)
	server.server_list = list
	server.server_using = ""
	server.conn = nil
	
	
	server.Redis_switch_server()

	if server.conn == nil {
		defaultServer = server
		e = NOVALIDSERVER
	} else {
		defaultServer = server
	}

	return
}

func (server * RedisToken) Redis_switch_server()  {
	
	if server.conn != nil {
		server.conn.Close()
		server.conn = nil
	}

	for _, v := range server.server_list {
			
		tmp := strings.Split(v,"@")
		host := tmp[0]
		var pass string
		if len(tmp) == 2 {
			pass = tmp[1]
		} else {
			pass = ""
		}
		
		conn, err := gore.Dial(host)

		if err == nil {
			err = conn.Auth(pass)
		} else {
			continue
		}

		if err == nil {
			server.server_using = v
			server.conn = conn
			break
		} 
	}
}

func (server * RedisToken) Redis_server_set(key string, val string, tm int) (e error) {

	if server.conn == nil {
		return NOVALIDSERVER
	}

	rep, e := gore.NewCommand("SET", key, val).Run(server.conn)

	if e == nil {
		if rep.IsError() {
			ret, _ := rep.Error()
			e = errors.New(ret)
			return
		}
	} 


	if e == nil {
		_, e = gore.NewCommand("EXPIRE", key, tm).Run(server.conn)
	}
 
	return
}

func (server * RedisToken) Redis_server_get(key string) (ret string, e error) {
	if server.conn == nil {
		ret = ""
		e = NOVALIDSERVER
		return 
	}

	rep, e := gore.NewCommand("GET", key).Run(server.conn)
	if e != nil {
		return
	}

	if rep.IsError() {
		ret, _ = rep.Error()
		e = errors.New(ret)
		return
	}

	ret, _ = rep.String()
	return
}

func (server * RedisToken) Redis_server_del(key string) (ret string, e error) {
	if server.conn == nil {
		ret = ""
		e = NOVALIDSERVER
		return 
	}

	rep, e := gore.NewCommand("DEL", key).Run(server.conn)
	if e != nil {
		return
	}

	if rep.IsError() {
		ret, _ = rep.Error()
		e = errors.New(ret)
		return
	}

	ret, _ = rep.String()
	return
}

/*
func (server * RedisToken)  Redis_server_hset(key string, index string, val string) (e error) {
	if server.conn == nil {
		return NOVALIDSERVER
	}
	_, e = gore.NewCommand("HSET", key, index, val).Run(server.conn)
	if e == nil {
		_, e = gore.NewCommand("EXPIRE", key, 120).Run(server.conn)
	}
	return
}

func  Redis_server_hset(key string, index string, val string) (e error) {

	e = defaultServer.Redis_server_hset(key, index, val)
	
	if e != nil {
		defaultServer.Redis_switch_server()
	}else{
		return
	}
	
	if defaultServer.conn != nil {
		return defaultServer.Redis_server_hset(key, index, val)
	} else {
		return
	}

}

*/
