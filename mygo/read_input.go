package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please input some word:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input is: %s \n", input)
	}
}
