package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var utf8Str []rune

	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[size:]
		utf8Str = append(utf8Str, r)
	}

	fmt.Println(utf8Str)

	str := string(utf8Str)
	bytesSend := []byte(str)

	fmt.Println(bytesSend)

	i16 := 0xff
	fmt.Printf(" i16-->i10 %d \n", i16)
}
