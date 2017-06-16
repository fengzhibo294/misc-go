package main

import (
	"fmt"
)

func main() {

	slice := []byte(" Google")

	fmt.Printf("before slice: %s \n", slice)
	ilen := len(slice) - 1
	for i, j := 0, ilen; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	fmt.Printf("after reverse slice : %s", slice)
}
