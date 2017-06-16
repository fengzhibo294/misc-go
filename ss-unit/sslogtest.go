package main

import (
	"sslog"
	//"fmt"
	//"os"
)

func main() {
	a := 4
	b := 5
	sslog.LoggerInit()
	sslog.LoggerDebug("init test 111 %d", a)
	sslog.LoggerErr("haaaaaaaaaa %d", b)
}
