package main

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
)

func test1() {
	fmt.Print("test")
	l4g.AddFilter("stdout", l4g.ERROR, l4g.NewConsoleLogWriter())
	l4g.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter("test.log", false))
	l4g.Debug("the time is now :%s -- %s", "213", "sad")
	l4g.Error("ttttttttttttttttttt")
	l4g.Info("testtt")
	l4g.Info("111111111")
	defer l4g.Close()
}

func test2() {
	log := l4g.NewLogger()

	// Create a default logger that is logging messages of FINE or higher
	log.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter("test2.log", false))
	log.Error("twoo")
	defer l4g.Close()
}

func test3() {
	l4g.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter("test.log", false))
	l4g.Error("test3")
	defer l4g.Close()

}

func main() {
	test1()
	test3()
	//test2()

}
