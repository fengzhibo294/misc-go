package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int
	Sex  int
}

func main() {
	p := Person{Name: "zhang", Age: 20, Sex: 1}
	j, _ := json.Marshal(p)
	fmt.Println(string(j))
}
