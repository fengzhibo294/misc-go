package main

import (
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"strconv"
)

type TestDB struct {
	a    int
	bStr string
}

func main() {
	newJs, _ := sjson.NewJson([]byte(`{}`))

	arrDB := []TestDB{}
	for i := 0; i < 5; i++ {
		bStr := "hello " + strconv.Itoa(i)
		uDB := &TestDB{i, bStr}
		uJs, _ := json.Marshal(uDB)

	}
	newJs.Set("test", uJs)
	printJs, _ := newJs.MarshalJSON()
	fmt.Printf("printJs json string:%s\n", printJs)
}
