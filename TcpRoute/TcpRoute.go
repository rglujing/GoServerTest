package TcpRoute;

import (
	"net"
	"encoding/binary"
	"bytes"
	l4g "code.google.com/p/log4go"
	"TokenServ/crypt"
	"TokenServ/token"
	)

const (
	ASHEADLEN = 16
)

const (
	/*创建授权码*/
 	AU_CREATE_AUCODE_REQ  = 0x3C01
 	AU_CREATE_AUCODE_RESP = 0x3C02

	/*获取Token */
	AU_REQUEST_TOKEN_REQ  = 0x3C03
	AU_REQUEST_TOKEN_RESP = 0x3c04

	/*访问AsToken*/
	AU_ACCESS_TOKEN_REQ   = 0x3C05
	AU_ACCESS_TOKEN_RESP  = 0x3C06

	/*刷新AsToken*/
	AU_REFRESH_TOKEN_REQ  = 0x3C07
	AU_REFRESH_TOKEN_RESP = 0x3C08

	/*检查AsToken*/
	AU_INSPECT_TOKEN_REQ  = 0x3C09
	AU_INSPECT_TOKEN_RESP = 0x3C0A

	/*注销Token*/
	AU_DESTROY_TOKEN_REQ   = 0x3C0B
	AU_DESTROY_TOKEN_RESP  = 0x3C0C

	/*申请Session*/
	AU_REQUEST_SESSION_REQ  = 0x3C0D
	AU_REQUEST_SESSION_RESP = 0x3C0E

	/*写入session数据*/
	AU_SET_SESSION_DATA_REQ = 0x3C13
	AU_SET_SESSION_DATA_RESP = 0x3C14
)


type Tcpheader struct {
	HeadID     uint16 
	Magic      byte 
	Version    uint8
	BodyLen    uint32
	VisitorIP  uint32
	Method     uint16
	PrivLen    uint16
}


func Token_tcp_listen(port string)  {

	ln, e := net.Listen("tcp", port)
	
	l4g.Debug("TCP listen at %s", port)

	if e != nil {
		l4g.Error("tcp listen error %s", e.Error())
		return
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleTcpConn(conn)
	}
}


func handleTcpConn(conn net.Conn) {
	
	var buf_header [ASHEADLEN]byte
	var buf_cont []byte
	var head Tcpheader
	/*ret is json str*/
	var ret string
	var para string

	defer conn.Close()

	addr := conn.RemoteAddr().String()


	/*Parse Header*/
	length, e := conn.Read(buf_header[:])
	if e != nil {
		l4g.Error(e.Error())
		return
	}

	if length != ASHEADLEN {
		l4g.Error("Receive uncomplete package : %s", addr)
		return
	}
	
	io_header  := bytes.NewBuffer(buf_header[:])
	e = binary.Read( io_header , binary.LittleEndian, &head)
	if e != nil {
		l4g.Error("invalid header %s %s", addr, e.Error())
		ret = token.Invalid_method_json
		head.Method = head.Method+1
		goto ErrorLoop
	}
	

	/*Read Contents*/
	/*Check body length, if too max, then return error */
	if (token.MAX_JSON_LEN != 0) && (head.BodyLen > uint32(token.MAX_JSON_LEN)) {
		ret = token.Conten_tlong_json
		head.Method = head.Method+1
		goto ErrorLoop
	}
	
	buf_cont = make([]byte, head.BodyLen)
	conn.Read(buf_cont)
	para = tokenaes.Decrypt(buf_cont)
	
	/*Do business*/

	switch head.Method {
		case  AU_CREATE_AUCODE_REQ :
			ret = token.Entry_Filter(para, addr, token.CREATE_AUCODE)
			head.Method = AU_CREATE_AUCODE_RESP
		case  AU_REQUEST_TOKEN_REQ :
			ret = token.Entry_Filter(para, addr, token.REQUEST_TOKEN)
			head.Method = AU_REQUEST_TOKEN_RESP
		case  AU_ACCESS_TOKEN_REQ :
			ret = token.Entry_Filter(para, addr, token.ACCESS_TOKEN)
			head.Method = AU_ACCESS_TOKEN_RESP
		case  AU_REFRESH_TOKEN_REQ :
			ret = token.Entry_Filter(para, addr, token.REFRESH_TOKEN)
			head.Method = AU_REFRESH_TOKEN_RESP
		case  AU_INSPECT_TOKEN_REQ :
			ret = token.Entry_Filter(para, addr, token.INSPECT_TOKEN)
			head.Method = AU_INSPECT_TOKEN_RESP
		case  AU_DESTROY_TOKEN_REQ :
			ret = token.Entry_Filter(para, addr, token.DESTROY_TOKEN)
			head.Method = AU_DESTROY_TOKEN_RESP
		case  AU_REQUEST_SESSION_REQ : 
			ret = token.Entry_Filter(para, addr, token.REQUEST_SESSION)
			head.Method = AU_REQUEST_SESSION_RESP
		case  AU_SET_SESSION_DATA_REQ :
			ret = token.Entry_Filter(para, addr, token.SET_SESSION_DATA)
			head.Method = AU_SET_SESSION_DATA_RESP
		default :
			ret = token.Invalid_method_json
	}

ErrorLoop:
	Result := tokenaes.Encrypt(ret)
	head.BodyLen = uint32(len(Result))

	err := binary.Write(io_header, binary.LittleEndian, head)
        if err != nil {
                l4g.Error("binary.Write failed:", err)
		return
        } 
	/*Return Result*/
	conn.Write( io_header.Bytes() )
	conn.Write( Result )

	return
	

}
