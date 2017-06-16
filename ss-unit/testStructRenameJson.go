package main

import (
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"strconv"
)

type Server struct {
	ServerName string //ServerName 首字母必须大写
	ServerIP   int
}

type renameServerslice struct {
	UserDBs []Server `json:"map"` //编码成json的时候,"UserDBs" ---> "map"
}

type Serverslice struct {
	UserDBs []Server
}

func ServersliceNomal() {
	var s Serverslice

	fmt.Println("Enter  ServersliceNomal() printf: ")
	s.UserDBs = append(s.UserDBs, Server{ServerName: "Shanghai_VPN", ServerIP: 123})
	s.UserDBs = append(s.UserDBs, Server{ServerName: "Beijing_VPN", ServerIP: 456})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
func main() {
	var s renameServerslice
	s.UserDBs = append(s.UserDBs, Server{ServerName: "Shanghai_VPN", ServerIP: 123})
	s.UserDBs = append(s.UserDBs, Server{ServerName: "Beijing_VPN", ServerIP: 456})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

	newJs, _ := sjson.NewJson(b)

	jsByte, _ := newJs.MarshalJSON()
	fmt.Println(string(jsByte))
	fmt.Println(strconv.FormatInt(int64(0x789), 16))
	//nomal struct
	ServersliceNomal()
}
