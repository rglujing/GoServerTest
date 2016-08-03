package tokenbdb;

import (
	"github.com/keimoon/gore"
	"errors"
	"strings"
	)

type BdbToken struct {
	server_list []string
	server_using string
	conn *gore.Conn
	e error
}

var defaultserver *BdbToken

var (
        NOVALIDSERVER = errors.New("invalid bdb server list")
)


func Bdb_server_set(key string, val string) (err error) {
	
	err = defaultserver.Bdb_server_set(key, val)

	if err != nil {
		defaultserver.Bdb_server_switch()
	} else {
		return
	}

	if defaultserver.conn != nil {
		return defaultserver.Bdb_server_set(key, val)
	} else {
		return
	}
}

func Bdb_server_get(key string) (val string, err error) {

	val, err = defaultserver.Bdb_server_get(key)

	if err != nil {
		defaultserver.Bdb_server_switch()
	} else {
		return
	}
	
	if defaultserver.conn != nil {
		return defaultserver.Bdb_server_get(key)
	} else {
		return
	}
}

func Bdb_server_del(key string) (err error) {

	err = defaultserver.Bdb_server_del(key)

	if err != nil {
		defaultserver.Bdb_server_switch()
	} else {
		return
	}
	
	if defaultserver.conn != nil {
		return defaultserver.Bdb_server_del(key)
	} else {
		return
	}
}

func Bdb_init( list []string ) ( server * BdbToken , e error ) {
	
	server = new(BdbToken)
	server.server_list = list
	server.server_using = ""
	server.conn = nil
	
	server.Bdb_server_switch()

	if server.conn == nil {
		defaultserver = server
		e = NOVALIDSERVER
	} else {
		defaultserver = server
	} 

	return
}

func (server * BdbToken) Bdb_server_switch() {
	
	if server.conn != nil {
		server.conn.Close()
		server.conn = nil
	}
	
	for _,v := range server.server_list {
		tmp := strings.Split(v,"@")
		host := tmp[0]
		var pass string
		if  len(tmp) == 2 {
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

func (server * BdbToken) Bdb_server_set(key string, val string) (err error){

	if server.conn == nil {
		return NOVALIDSERVER
	}

	rep,err := gore.NewCommand("SET", key, val).Run(server.conn)

	if err == nil {
		
		if rep.IsError() {
			ret, _ := rep.Error()
			err = errors.New(ret)
			return
		}
	
	}
	return
}

func (server * BdbToken) Bdb_server_get(key string) (ret string, err error) {

	if server.conn == nil {
		ret = ""
		err = NOVALIDSERVER
		return
	}

	rep, err := gore.NewCommand("GET", key).Run(server.conn)
	if err != nil {
		return
	}


	if rep.IsError() {
		ret, _ = rep.Error()
		err = errors.New(ret)

		return
	}

	ret, _ = rep.String()
	return

}

func (server * BdbToken) Bdb_server_del(key string) (err error) {

	if server.conn == nil {
		err = NOVALIDSERVER
		return
	}

	rep, err := gore.NewCommand("DEL", key).Run(server.conn)
	if err != nil {
		return
	}

	if rep.IsError() {
		ret, _ := rep.Error()
		err = errors.New(ret)
		return
	}

	return
}
