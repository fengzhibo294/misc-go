package main

import (
	"fmt"
	"time"
)

func closeOutTime() {
	fmt.Println("test ")
}

func main() {
	/*
		st := "2013-10-28 13:20:11 (CEST)"
		loc, _ := time.LoadLocation("Local")
		time2, err := time.ParseInLocation("2006-12-31 23:59:59", st, loc)
		fmt.Println(time2)
		fmt.Println(err)
	*/
	time2, err := time.Parse("2006-01-02 15:04:05", "2013-11-18 07:52:00")
	fmt.Println(time2)
	fmt.Println(err)

	//go closeOutTime()
	go func() {
		for {
			closeOutTime()
			time.Sleep(1e9)
		}
	}()

	time.Sleep(10 * 1e9)
}
