package main

import (
	"fmt"
)

func main() {

	arr := [5]int{5, 1, 4, 2, 8}

	fmt.Println("array is : ", arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	fmt.Println("array --->>bubble: ", arr)
}
