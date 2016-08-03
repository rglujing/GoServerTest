package main;

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"strings"
	)   


func httpRequest(jsonstr string, pad string)  {

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify:true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	var res *http.Response
	var err error
	/*
	if get {
		res, err = client.Get("https://localhost:8188/AccService/OAuth2/" + pad + "?" + jsonstr)
	} else {
		res, err = client.Post("https://localhost:8188/AccService/OAuth2/" + pad, "application/json", strings.NewReader( jsonstr ) )
	}
	*/
	res, err = client.Post(conf.Http + pad, "application/json", strings.NewReader( jsonstr ) )
	if err != nil {
		fmt.Printf("err is %s\n", err.Error())
		return
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Printf("err is %s\n", err.Error())
		return
	}

	fmt.Printf("result is %s\n", string(robots))
}
