package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	for {
		i := <-in
		fmt.Println(i)
	}

}

func f(out chan int) {
	out <- 2
	//out <- 3
}

func main() {
	out := make(chan int)

	go f(out)
	go f1(out)

	time.Sleep(1 * 1e9)
}
