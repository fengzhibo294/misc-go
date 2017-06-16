package main

import (
	"fmt"
	"os"
	//"strconv"
)

/*
func main() {
	arg_num := len(os.Args)
	fmt.Printf("the num of input is %d\n", arg_num)

	fmt.Printf("they are :\n")
	for i := 0; i < arg_num; i++ {
		fmt.Println(os.Args[i])
	}

	sum := 0
	for i := 1; i < arg_num; i++ {
		curr, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println("error happened ,exit")
			os.Exit(1)
		}
		sum += curr
	}

	fmt.Printf("sum of Args is %d\n", sum)
	os.Exit(0)
}
*/
func main() {
	arg_num := len(os.Args)
	if arg_num <= 1 {
		fmt.Println("not enough args: " + os.Args[0])
		os.Exit(1)
	}

	str := os.Args[1]
	fmt.Printf("port:%s \n", str)
	os.Exit(0)
}
