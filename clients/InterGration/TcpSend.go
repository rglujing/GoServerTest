package main;

import (
	"net"
	"fmt"
	"bytes"
	"encoding/binary"
	"crypto/aes"
	"strings"
	)

type header struct {
	HeadID     uint16 
	Magic      byte 
	Version    uint8
	BodyLen    uint32
	VisitorIP  uint32
	Method     uint16
	PrivLen    uint16
}

const (
	ASHEADLEN = 16
	AU_CREATE_AUCODE_REQ = 0x3C01
	AU_REQUEST_TOKEN_REQ  = 0x3C03
	AU_ACCESS_TOKEN_REQ   = 0x3C05
	AU_REFRESH_TOKEN_REQ  = 0x3C07
	AU_INSPECT_TOKEN_REQ  = 0x3C09
	AU_DESTROY_TOKEN_REQ   = 0x3C0B
	AU_REQUEST_SESSION_REQ  = 0x3C0D
	AU_SET_SESSION_DATA_REQ = 0x3C13
)

var key []byte = []byte{
        1, 2, 3, 4, 5, 6, 7, 8,
        9, 0, 1, 2, 3, 4, 5, 6,
        7, 8, 9, 0, 1, 2, 3, 4,
        5, 6, 7, 8, 9, 0, 1, 2,
}


func  SendMsg(buf *bytes.Buffer) {
	
	conn, err := net.Dial("tcp", conf.Tcp)
	if err != nil {
		fmt.Printf("conn err is %s\n", err.Error())
		return
	}
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("send err is %s\n", err.Error())
		return
	}

	var buf_header [ASHEADLEN]byte
	var buf_cont []byte
	var head header

	conn.Read(buf_header[:])
	io_header  := bytes.NewBuffer(buf_header[:])
	e := binary.Read( io_header , binary.LittleEndian, &head)
        if e != nil {
		fmt.Printf("invalid header %s", e.Error())
        }

	buf_cont = make([]byte, head.BodyLen)
	conn.Read(buf_cont)
	ret := Decrypt(buf_cont) 

	fmt.Printf("result is %s\n", ret)
}


func Encrypt(src string) (dst []byte) {

        cleartext := make([]byte, aes.BlockSize)
        ciphertext := make([]byte, aes.BlockSize)

        cip, _ := aes.NewCipher(key)
        tmpReader := strings.NewReader(src)

        for _, e := tmpReader.Read(cleartext); e == nil; _, e = tmpReader.Read(cleartext) {
                cip.Encrypt(ciphertext, cleartext)
                dst = append(dst, ciphertext...)
                cleartext = make([]byte, aes.BlockSize)
                ciphertext = make([]byte, aes.BlockSize)
        }   
        return
}



func Decrypt(src []byte) (dst string) {

        cleartext := make([]byte, aes.BlockSize)
        ciphertext := make([]byte, aes.BlockSize)

        cip, _ := aes.NewCipher(key)
        tmpReader := bytes.NewReader(src)
        var rst []byte
        for _, e := tmpReader.Read(ciphertext); e == nil; _, e = tmpReader.Read(ciphertext) {
                cip.Decrypt(cleartext, ciphertext)
                rst = append(rst, cleartext...)
                cleartext = make([]byte, aes.BlockSize)
                ciphertext = make([]byte, aes.BlockSize)
        }   
    
        return string(rst)
}


func tcpRequest(cont string, METH uint16) {
	
        aesstr := Encrypt(cont)
	head := header{
			HeadID:0x0ABA,
			Magic:byte('$'), 
			Version:0x30,
			BodyLen:uint32(len(aesstr)),
			VisitorIP:0,
			Method: METH,
			PrivLen:0}


	buf := new(bytes.Buffer)
	
	err := binary.Write(buf, binary.LittleEndian, head)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

        buf.Write(aesstr)
	
	SendMsg(buf)
}


