package main

import (
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"strings"
	"time"
)

func main() {
	var ch chan interface{}

	TestSetJson()
	TestSetPath()
	TestSetPath2()
	ch = make(chan interface{}, 50)
	go doSend(ch)
	go doRecv(ch)

	time.Sleep(10 * 1e9)
}

func TestSetPath() {
	js, _ := sjson.NewJson([]byte(`{}`))

	js.SetPath([]string{"foo", "bar"}, "baz") //&{map[foo:map[bar:baz]]}
	s, _ := js.GetPath("foo", "bar").String() //baz
	fmt.Println("TestSetPath: ", s)
	fmt.Println(js)
}

func TestSetPath2() {
	js, _ := sjson.NewJson([]byte(`{}`))

	js.SetPath([]string{"sub_obj", "a"}, 1) //&{map[sub_obj:map[a:1]]}
	js.SetPath([]string{"sub_obj", "b"}, 5) //&{map[sub_obj:map[b:5]]}

	s, _ := js.GetPath("sub_obj", "a").Int() //1
	sub := js.Get("sub_obj")
	fmt.Println("TestSetPath2: ", s)
	fmt.Println(js)
	fmt.Println("sub: ", sub) // &{map[a:1 b:5]}
	subJson, _ := sub.MarshalJSON()
	fmt.Printf("sub json:%s", subJson)
	fmt.Printf("sub json string:%s\n", subJson)

	strJS, _ := js.MarshalJSON()
	fmt.Printf("js json string:%s\n", strJS)
}

func TestSetJson() {
	bTest := false
	fmt.Printf("test strconv bool: %s \n", bTest)
	js, _ := sjson.NewJson([]byte(`{}`))
	js.Set("a", 123)
	js.Set("b", 456)
	//浏览同步通知
	informJs, _ := sjson.NewJson([]byte(`{}`))
	informJs.Set("cid", 12) //cid -1单发回复
	informJs.SetPath([]string{"sendbuf", "m"}, 1032)
	//var jser interface{}
	//jser = js
	informJs.SetPath([]string{"sendbuf", "p"}, js)

	strJs, _ := informJs.MarshalJSON()
	fmt.Printf("TestSetJson informJs string: %s \n", strJs)
}

func doSend(ch chan interface{}) {

	var djs *sjson.Json
	var err error
	djs, err = sjson.NewJson([]byte(`{}`))
	if err != nil {
		fmt.Println(err)
	}

	var ts interface{}
	ts = djs

	fmt.Println(djs)
	var njs *sjson.Json
	njs = ts.(*sjson.Json)
	njs.Set("baz", "bing")
	njs.Set("cid", 5)
	//njs.Set("sub_obj").Set("a", 4)
	//njs.Set("sub_obj").Set("b", 6)
	//subJs := njs.Get("sub_obj")
	//fmt.Println("sub_obj", subJs)
	s, _ := njs.Get("baz").String()
	i, _ := njs.Get("cid").Int()
	fmt.Printf("send baz: %s cid[%d]\n", s, i)
	ch <- *njs //通道
	ch <- "test"

}

func doRecv(ch chan interface{}) {
	for {
		//var js *sjson.Json
		in := <-ch
		switch vv := in.(type) {
		case string:
			fmt.Println("string", vv)
			bTrue := strings.Contains(vv, "te")
			fmt.Println("contain: ", bTrue)
		case sjson.Json:
			fmt.Println("sjson.Json", vv)
		}
		/*
			js := in.(*sjson.Json)

			fmt.Println("js")
			s, _ := js.Get("baz").String()
			fmt.Printf("recv baz: %s \n", s)*/
	}

}
