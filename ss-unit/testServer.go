package main

import (
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"net"
	"runtime"
	"strings"
)

func WebSocketEncode(inBytes []byte, session *config.TcpSession) []byte {
	var outBytes []byte
	outLen := 0

	//对websocket加上头部字段，一条完整的文本消息
	outBytes = append(outBytes, byte(129)) // 详细参考http://blog.csdn.net/otypedef/article/details/51492188

	//对websocket 加上长度字段
	inLen := len(inBytes)
	if inLen <= 125 { //占1字节长度
		outLen = inLen
		outBytes = append(outBytes, byte(outLen))
	} else if inLen <= 65535 { //占3字节长度
		outLen = outLen + 126
		outBytes = append(outBytes, byte(outLen))
		outBytes = append(outBytes, byte(inLen>>8))
		outBytes = append(outBytes, byte(inLen))
	} else { //>65535 占8个字节长度
		outLen = outLen + 127
		outBytes = append(outBytes, byte(outLen))

		outBytes = append(outBytes, byte(0))
		outBytes = append(outBytes, byte(0))
		outBytes = append(outBytes, byte(0))
		outBytes = append(outBytes, byte(0))

		outBytes = append(outBytes, byte(inLen>>24))
		outBytes = append(outBytes, byte(inLen>>16))
		outBytes = append(outBytes, byte(inLen>>8))
		outBytes = append(outBytes, byte(inLen))
	}

	outBytes = append(outBytes, inBytes...)
	return outBytes
}

func WebSocketDecode(inBytes []byte) *sjson.Json {
	inLen := len(inBytes)
	if inLen < 2 {
		return nil
	}

	wsHeadflag := inBytes[0]   //web socket 1个字节头部标示
	wsFin := (wsHeadflag >> 7) //是否结尾包
	wsOp := wsHeadflag & 0xF   //消息类型

	if wsFin != 1 {
		fmt.Println("Decoder : Need Buffer")
		return nil
	}

	wsIsMsgData := false
	switch wsOp {
	case 0:
		fmt.Println("Decoder : Frame Data")
	case 1, 2:
		wsIsMsgData = true
	case 8, 9, 10:
		wsIsMsgData = false
	default:
		return nil
	}

	wsHeadpayload := inBytes[1] //websocket头部内容长度属性
	payloadlen := wsHeadpayload & 0x7F
	mask := wsHeadpayload >> 7
	headlength := 2 //至少2个字节的websocket头部长度

	//无数据处理
	if wsIsMsgData != true {
		return nil
	}

	//有数据需要处理, 定位到json数据开始的位置
	if payloadlen == 126 { //0x7E
		headlength = headlength + 2
	} else if payloadlen == 127 { //0x7F
		headlength = headlength + 8
	}

	if mask == 1 {
		headlength = headlength + 4
	}

	outJson, err := sjson.NewJson(inBytes[headlength:])
	if err != nil {
		fmt.Println("recv json body not correct --------->>>headlength:", headlength, "inBytes[headlength:]", string(inBytes[headlength:]))
		fmt.Println("inBytes:", string(inBytes), "inBytes[headlength]: ", inBytes[headlength])
		fmt.Println("err:", err)
		return nil
	}

	return outJson
}

func main() {
	service := ":80"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Printf("ResolveTCPAddr: %s \n", err.Error())
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("ListenTCP: addr[ip:port]  %s\n", err.Error())
	}

	fmt.Printf("ListenTCP: %s success \n\n", service)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("local addr:", conn.LocalAddr(), "tcpAddr.Port:", tcpAddr.Port)
		fmt.Println("a client accept: ", conn.RemoteAddr())
		go DoRecvDecode(conn)
	}
}

func DoRecv(conn net.Conn) {
	recvbuf := make([]byte, 10240)

	for {
		//读取数据
		iLen, err := conn.Read(recvbuf)
		if err != nil {

			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("[%s,%d]doRecv--recv failed to close  err:%v \n", file, line+1, err)
			break
		}

		_, file, line, _ := runtime.Caller(0)
		fmt.Printf("[%s,%d]doRecv 接收到的数据长度 iLen =%d \n", file, line+1, iLen)
		if iLen < 0 {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("[%s:%d]doRecv--iLen < 0 continue", file, line+1)
			break
		} else {
			readStr := string(recvbuf[0:iLen])
			if bKey := strings.Contains(readStr, "LkCwIf8QOUGUOUajbgIIqA=="); bKey {
				//fmt.Println("recvBuf :", readStr)
				fmt.Println("LkCwIf8QOUGUOUajbgIIqA== contains : ", bKey)
				ResponseHandShake(conn)
			} else {
				firstBytes := []byte{129, 254, 01, 149, 0, 0, 0, 0}
				firstStr := string(firstBytes)

				trimString := strings.TrimLeft(readStr, firstStr)
				fmt.Printf("recv json : %s \n", trimString)
				//读取的数据处理
				js, err := sjson.NewJson([]byte(trimString))
				if err != nil || js == nil { //处理出错，json为空
					_, file, line, _ := runtime.Caller(0)
					fmt.Printf("[%s:%d]doRecv--err != nil || js == nil  recvbuf: %v: len[%d] \n", file, line+1, string(recvbuf), len(recvbuf[0:iLen]))
					continue
				} else {
					fmt.Printf("it is json data \n")
				}
			}
		}

	}

	fmt.Println("exit DoRecv() break")
	conn.Close()
	return
}

func ResponseHandShake(conn net.Conn) {
	serverKey := "WRVoTRt3LIcNYP32yfQM44zmAkY="

	httpStr := "HTTP/1.1 101 Web Socket Protocol Handshake\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: " + serverKey + "\r\n\r\n"
	writeLen, _ := conn.Write([]byte(httpStr))
	fmt.Println("writeLen : ", writeLen)
}

func DoRecvDecode(conn net.Conn) {
	recvbuf := make([]byte, 10240)

	for {
		//读取数据
		iLen, err := conn.Read(recvbuf)
		if err != nil {

			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("[%s,%d]doRecv--recv failed to close  err:%v \n", file, line+1, err)
			break
		}

		_, file, line, _ := runtime.Caller(0)
		fmt.Printf("[%s,%d]doRecv 接收到的数据长度 iLen =%d \n", file, line+1, iLen)
		if iLen < 0 {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("[%s:%d]doRecv--iLen < 0 continue", file, line+1)
			break
		} else {
			readStr := string(recvbuf[0:iLen])
			if bKey := strings.Contains(readStr, "LkCwIf8QOUGUOUajbgIIqA=="); bKey {
				//fmt.Println("recvBuf :", readStr)
				fmt.Println("LkCwIf8QOUGUOUajbgIIqA== contains : ", bKey)
				ResponseHandShake(conn)
			} else {
				recvJson := WebSocketDecode(recvbuf[0:iLen])
				if recvJson == nil {
					fmt.Println("recvJson nil")
				} else {
					fmt.Println("recvJson :", recvJson)
				}
			}
		}

	}

	fmt.Println("exit DoRecv() break")
	conn.Close()
	return
}
