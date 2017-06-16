// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import (
	"fmt"
	"time"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; i < 10; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	fmt.Printf("filter Enter: \n\n")
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		fmt.Printf("i=%d, prime=%d \n", i, prime)
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
	fmt.Printf("filter Exit:\n\n")
}

// The prime sieve: Daisy-chain filter processes together.
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	time.Sleep(100 * 1e9)
}
