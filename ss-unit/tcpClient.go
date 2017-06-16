package main

import (
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"net"
	"os"
	"strconv"
	"time"
	//"unicode/utf8"
)

const (
	/**登陆登出模块*/
	LOGIN = 1
	/**列表模块*/
	LIST = 2
	/**聊天模块*/
	CHAT = 3
	/**权限模块*/
	POWER = 4
	/**上传模块*/
	UPLOAD = 5
	/**白板模块*/
	WHITEBOARD = 6
	/**视频模块*/
	VIDEO = 7
	/**浏览器模块*/
	BROWER = 8
)
const (
	Goroutings = 1
)

var (
	Gname     = "倪晓威"
	Guserid   = "159-12034123"
	RecvCount = 0
)

func main() {
	for i := 0; i < Goroutings; i++ {
		go tcpConnectClient(i)
	}

	for {
		time.Sleep(1e9)
	}
}

func tcpConnectClient(personID int) {
	service := "192.168.70.182:80"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	fmt.Printf("connect successful  remote addr: %s \n", conn.RemoteAddr().String())

	//发送json数据
	sendJsonData(conn, personID)
	for {
		//result, err := readFully(conn)
		var readBuf [1024]byte
		readLen, err := conn.Read(readBuf[0:])
		checkError(err)

		RecvCount++
		fmt.Printf("RecvCount=%d [time:%v]recv data:%s \n", RecvCount, time.Now(), string(readBuf[0:readLen]))

		//fmt.Println("exit main")
		//os.Exit(0)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s \n", err.Error())
		os.Exit(1)
	}
}

func newPjson(personID int) *sjson.Json {
	var userid, name string
	if personID == 0 {
		userid = Guserid
		name = Gname
	} else {
		userid = Guserid + strconv.Itoa(personID)
		name = Gname + strconv.Itoa(personID)
	}

	pJs, _ := sjson.NewJson([]byte(`{}`))
	pJs.Set("cid", 1000001)
	pJs.Set("userid", userid)
	pJs.Set("name", name)
	pJs.Set("character", 1)

	return pJs
}

func sendJsonData(conn *net.TCPConn, personID int) error {
	const CLIENT_LOGIN = 1

	//生成发送的json数据
	sendJs, _ := sjson.NewJson([]byte(`{}`))
	sendJs.Set("m", LOGIN*1000+CLIENT_LOGIN)
	sendJs.Set("p", newPjson(personID))
	fmt.Printf("sendJs=%v \n", sendJs)

	//转换成[]byte格式数据进行发送
	sendBuf, _ := sendJs.MarshalJSON()

	WriteRune(conn, sendBuf)
	/*
		//转换成要发送的切片
		utf8Str := string(sendRune)
		sendBytes := []byte(utf8Str)

		fmt.Printf("[time:%v]sendBuf=%s %v\n", time.Now(), sendStr, sendBytes)
		_, err := conn.Write(sendBuf)
		//_, err = conn.Write([]byte("\n"))
		checkError(err)
	*/
	return nil
}

func WriteRune(conn *net.TCPConn, sendBytes []byte) {
	//const BOM = '\ufffe'
	//var bytes [4]byte
	//var sendBuf []byte
	/*
		bytes[0] = BOM >> 8
		bytes[1] = BOM & 255
		_, err := conn.Write(bytes[0:2])
		if err != nil {

			return

		}
	*/

	/*
		count := 0
		for len(sendBytes) > 0 {
			r, size := utf8.DecodeRune(sendBytes)

			bytes[0] = byte(r >> 24)
			bytes[1] = byte(r >> 16)
			bytes[2] = byte(r >> 8)
			bytes[3] = byte(r & 255)
			sendBytes = sendBytes[size:]
			i := 0
			for i = 0; i < size; i++ {
				sendBuf = append(sendBuf, bytes[i])
			}

			count = count + size
		}

		fmt.Println("count: ", count)
		sendLen, err := conn.Write(sendBuf[0:count])
		if err != nil {

			return

		}
	*/
	sendLen, err := conn.Write(sendBytes[0:])
	if err != nil {

		return

	}

	fmt.Printf("WriteRune success sendLen = %d\n", sendLen)
	return
}
