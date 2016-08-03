package main;

import (
	"github.com/coreos/go-etcd/etcd"
	"fmt"
	"strings"
	"time"
	)


type Etcd_token struct {
	cli   *etcd.Client
	index uint64         // for watch function, since index
	NodeName string
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
	
	defaultet = et
	return	
}


func Etcd_Watch(key string) (e error) {
	return defaultet.Etcd_Watch(key)
}

var receiver chan *etcd.Response
var stop chan bool

func (et * Etcd_token) Etcd_Watch(key string) (e error) {
	
	receiver = make(chan *etcd.Response)
	stop = make(chan bool)
	
	go etcd_listener(et, receiver, stop)
	go et.cli.Watch(key, et.index, false, receiver, stop)

	return
}


func etcd_listener( et * Etcd_token, rec chan * etcd.Response, stop chan bool) {
	
	for {
		fmt.Printf("Listerner is working\n")
		t,ok := <-rec
		

		if t == nil {
			fmt.Printf("Why receive null\n")
			fmt.Printf("Listerner is stopping\n")
			fmt.Printf("response mask is %v\n", ok)

			//close(rec)
			close(stop)
			
			_,ok = <- stop
			fmt.Printf("stop mask is %v\n", ok)
			break
		} else {
		
			if t.Node != nil {
				fmt.Printf("t.Node.ModifiedIndex is %d\n", t.Node.ModifiedIndex)
				fmt.Printf("%s\n", t.Node.Value)
			} else {
				fmt.Printf("t.Node is nil")
			}

		}
		
	}
}

/*
func etcd_listener( et * Etcd_token, rec chan * etcd.Response, stop chan bool) {
	
	for {
		fmt.Printf("Listerner is working\n")
		t := <-rec
		
		if t == nil {
			fmt.Printf("Why receive null\n")
			fmt.Printf("Listerner is stopping\n")
			break
		} else {
		
			if t.Node != nil {
				fmt.Printf("t.Node.ModifiedIndex is %d\n", t.Node.ModifiedIndex)
				fmt.Printf("%s\n", t.Node.Value)
			} else {
				fmt.Printf("t.Node is nil")
			}

		}
		
	}


	fmt.Printf("Relistenning\n")
	et.Etcd_Watch("/tokenserv/tokendef")
}
*/

func main() {

	Etcd_init([]string{"https://10.15.201.55:4001", "https://10.15.201.55:4002"}, "../ca/etcd.pem", "../ca/etcd.key")
	Etcd_Watch("/tokenserv/tokendef")

	for {
		time.Sleep(time.Second)
	}
}
