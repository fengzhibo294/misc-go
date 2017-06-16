package main

import (
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   int
}

type Serverslice struct {
	UserDBs []Server
}

func main() {
	var s Serverslice
	s.UserDBs = append(s.UserDBs, Server{ServerName: "Shanghai_VPN", ServerIP: 123})
	s.UserDBs = append(s.UserDBs, Server{ServerName: "Beijing_VPN", ServerIP: 456})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

	newJs, _ := sjson.NewJson(b)

	arr, _ := newJs.Get("UserDBs").Array()
	myJs, _ := sjson.NewJson([]byte(`{}`))
	myJs.Set("map", myJs)
	printJs, _ := myJs.MarshalJSON()
	fmt.Println(string(printJs))
}
