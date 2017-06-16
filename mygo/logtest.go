package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func SystemMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}

	defer logfile.Close()
	t := time.Now()
	fmt.Println(t.Unix())
	fmt.Println(SystemMs() / 1000)
	sec := time.Second
	fmt.Printf("location second = %d\n", sec)
	loggerErr := log.New(logfile, "[error]: ", log.Ldate|log.Ltime)
	loggerDebug := log.New(logfile, "[debug]: ", log.Ldate|log.Ltime)
	loggerDebug.Print("hello")
	loggerDebug.Print("oh....")
	loggerErr.Print("test")
	loggerErr.Print("test2")

	log.SetOutput(logfile)
	log.SetPrefix("[error]: ")
	log.Print("test log")

}
