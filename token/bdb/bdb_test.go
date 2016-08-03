package tokenbdb;

import (
	"fmt"
	_"encoding/json"
	"testing"
	_"strconv"
	_"time"
)

func TestBDBSET(t *testing.T) {
	
	server, e := Bdb_init( []string{"127.0.0.1:6379", "10.15.201.55:42226"})
	
	if e != nil {
		t.Fatal(e)
	}
	

	server.Bdb_server_set("token:test", "{laugh}")
}

func TestBDBGET(t *testing.T) {
	
	server, e := Bdb_init( []string{"127.0.0.1:6379", "10.15.201.55:42226"})
	
	if e != nil {
		t.Fatal(e)
	}
	

	val,_ := server.Bdb_server_get( "token:test" )

	fmt.Println(val)
	
}
