package main

import (
	"fmt"
	"runtime"
)

func main() {
	//funcName, file, line, ok := runtime.Caller(0)
	if true {
		// /fmt.Println("Func Name=" + runtime.FuncForPC(funcName).Name())
		_, file, line, _ := runtime.Caller(0)
		fmt.Printf("file: %s line=%d\n", file, line+1)
	}
}
