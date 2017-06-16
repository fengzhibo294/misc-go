package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := strings.Split("1,a,b,c,d,e,", ",")
	fmt.Printf("%q\n", strings.Split(",a,b,c,d,e,", ","))
	fmt.Printf("%q\n", strings.Split(",a,b,c,d,e,", ""))
	fmt.Printf("%q\n", strings.Split(",a,b,-c,d,e,", ",-"))
	fmt.Printf("%q\n", strings.Split("Hello，world！中国，你好！", "！"))
	fmt.Printf("%q\n", strings.Split("Hello，world！中国，你好！", "，"))
	fmt.Printf("arr[0]: %s \n", arr[0])
} //会议结束 用户的离开时间
