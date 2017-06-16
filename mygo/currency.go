package main

import (
	"fmt"
	"time"
)

func doWork(i int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(i)
}

func main() {
	for i := 1; i <= 5; i++ {
		go doWork(i) // Concurrency!
	}

	time.Sleep(1 * 1e10)
}
