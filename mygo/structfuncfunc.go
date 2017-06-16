package main

import (
	"fmt"
)

type TestS struct {
	a int
	b int
	c int
	d int
}

func init() {
	fmt.Println("hello init2")
}

func init() {
	fmt.Println("hello init3")
}

func init() {
	fmt.Println("hello init0")
}

func init() {
	fmt.Println("hello init1")
}

func (this *TestS) Sum() int {
	return this.sum1() + this.Sum2()
}

func (this *TestS) sum1() int {
	return this.a + this.b
}

func (this *TestS) Sum2() int {
	return this.c + this.d
}

func main() {
	test := &TestS{1, 2, 4, 5}

	fmt.Printf("the sum:%d \n", test.Sum())
}
