package tokenredis;

import ( 
	"testing"
	"fmt"
	"time"
	)
/*
func TestRedisConn( t * testing.T) {
	
	s,e := Redis_init( []string{"127.0.0.1:6379", "127.0.0.1:6380"} )
	
	if e != nil {
		t.Fatal(e)
	}
	
	fmt.Printf("Using %s \n", s.server_using)
}
*/
func TestRedisSet( t * testing.T) {
	
	_,e := Redis_init( []string{"127.0.0.1:6379", "127.0.0.1:6380@123456"} )
	
	if e != nil {
		t.Fatal(e)
	}
	
	time.Sleep(time.Second*10)

	e = Redis_server_set("Testing", "hahahah")
	if e != nil {
		fmt.Printf("%s\n", e.Error())
		t.Fatal(e)
	}

	//fmt.Printf("Using %s %s\n", s.server_using)
}
/*
func TestRedisGet( t * testing.T) {
	
	s,e := Redis_init( []string{"127.0.0.1:6379", "127.0.0.1:6380"} )
	
	if e != nil {
		t.Fatal(e)
	}
	
	val, e := s.Redis_server_get("Testing")
	if e != nil {
		t.Fatal(e)
	}
	
	fmt.Printf("val is %s\n", val)
}
*/
