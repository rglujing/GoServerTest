package HttpClient;

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"strings"
	"crypto/aes"
	"bytes"
	)

func encrypt( src string ) ( dst []byte )  {
    key := []byte{
        1, 2, 3, 4, 5, 6, 7, 8,
        9, 0, 1, 2, 3, 4, 5, 6,
        7, 8, 9, 0, 1, 2, 3, 4,
        5, 6, 7, 8, 9, 0, 1, 2,
    }   
    
    cleartext := make([]byte, aes.BlockSize)
    ciphertext := make([]byte, aes.BlockSize)
	
    cip, _ := aes.NewCipher(key)

    tmpReader := strings.NewReader(src)

    for _,e := tmpReader.Read(cleartext); e == nil; _,e = tmpReader.Read(cleartext) {
    	cip.Encrypt(ciphertext, cleartext)
	dst = append(dst, ciphertext...)
	//fmt.Println(cleartext)
	//fmt.Println(ciphertext)
	cleartext = make([]byte, aes.BlockSize)
	ciphertext = make([]byte, aes.BlockSize)
    }
	
	return
}


func httpRequest(get bool, en bool, jsonstr string, pad string) (ret string) {


	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify:true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	var res *http.Response
	var err error
	if get {
		res, err = client.Get("https://localhost:8188/AccService/OAuth2/" + pad + "?" + jsonstr)
	} else if en {
		aesstr := encrypt(jsonstr)
		buf := new(bytes.Buffer)
		buf.Write([]byte(aesstr))
		res, err = client.Post("https://localhost:8188/AccService/OAuth2/" + pad, "application/encrypt", buf )
	} else {
		res, err = client.Post("https://localhost:8188/AccService/OAuth2/" + pad, "application/json", strings.NewReader( jsonstr ) )
	}


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
	fmt.Printf("Received : %s\n", robots)

	return string(robots)


}


func CreateAuCode(get bool, en bool, jsonstr string) (ret string){
	return httpRequest(get, en, jsonstr, "CreateAuCode")
}

func RequestToken(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "RequestToken")
}

func RequestSession(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "RequestSession")
}

func AccessToken(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "AccessToken")
}

func RefreshToken(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "RefreshToken")
}

func InspectToken(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "InspectToken")
}

func DestroyToken(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "DestroyToken")
}

func SetSessionData(get bool, en bool, jsonstr string) (ret string) {
	return httpRequest(get, en, jsonstr, "SetSessionData")
}
