package main

import (
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"strconv"
)

type UserDB struct {
	Video bool
	Id    int    `json:"id"`
	Name  string `json:"name"`
}
type UsersliceType struct {
	UserDBs []*UserDB
}

type UserDBArray struct {
	UserDBs []UserDB
}

func NormalValue() {
	var usrArray UserDBArray

	for i := 0; i < 5; i++ {
		strName := "helloValue" + strconv.Itoa(i)
		uDB := UserDB{i == 3, i, strName}
		usrArray.UserDBs = append(usrArray.UserDBs, uDB)
	}

	fmt.Println(usrArray)

	for i, value := range usrArray.UserDBs {
		if i == 0 {
			value.Name = "test pointer"
		}
	}

	fmt.Println(usrArray)
}

var (
	UserList    map[string]*UserDB
	SessionList map[int]map[string]*UserDB
)

func TestSessionList() {
	UserList = make(map[string]*UserDB)
	SessionList = make(map[int]map[string]*UserDB)

	tmpDB := UserDB{true, 456, "session test"}
	UserList["test0"] = &tmpDB
	SessionList[0] = UserList
	//SessionList[0]["test0"] = &UserDB{true, 456, "session test"}

	uDB := SessionList[0]["test0"]

	fmt.Printf("before --uDB %v \n", uDB)

	uDB.Video = false
	uDB.Id = 789
	uDB.Name = "rename test"

	fmt.Printf("SessionList[0][test0]       %v \n", SessionList[0]["test0"])
	fmt.Printf("tmpDB %v \n", tmpDB)

}

func TestPointer() *UserDB {
	strName := "hello" + strconv.Itoa(1)
	uDB := &UserDB{1 == 3, 1, strName}

	return uDB
}

func TestStringSlice() {
	var strSlice []string

	fmt.Println("TestStringSlice :")
	for i := 0; i < 4; i++ {
		iString := "test" + strconv.Itoa(i*100)
		strSlice = append(strSlice, iString)
	}

	fmt.Println(strSlice)

	b, err := json.Marshal(strSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("TestStringSlice b json string:%s \n", string(b))
	}

	newJs, _ := sjson.NewJson(b)
	fmt.Printf("TestStringSlice newJs: %v \n", newJs)

	strJs, _ := sjson.NewJson([]byte(`{}`))
	strJs.Set("audiolist", newJs)

	strBytes, _ := strJs.MarshalJSON()
	fmt.Println("TestStringSlice strJs", string(strBytes))
}

func main() {
	var usrSlice UsersliceType

	TestSessionList()
	for i := 0; i < 5; i++ {
		strName := "hello" + strconv.Itoa(i)
		uDB := &UserDB{i == 3, i, strName}
		usrSlice.UserDBs = append(usrSlice.UserDBs, uDB)
	}

	test := *TestPointer()
	fmt.Printf("%v \n", test)
	fmt.Println(usrSlice)
	for i, value := range usrSlice.UserDBs {
		if i == 0 {
			value.Name = "test pointer"
		}
	}

	b, err := json.Marshal(usrSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("b json string:%s \n", string(b))
	}

	newJs, _ := sjson.NewJson(b)
	fmt.Printf("newJs: %v \n", newJs)

	sendJs, _ := sjson.NewJson([]byte(`{}`))
	sendJs.Set("p", newJs)

	sendBuf, _ := sendJs.Get("p").MarshalJSON()
	fmt.Println("sendBuf: " + string(sendBuf))

	fmt.Println("\n\n")
	NormalValue()
	TestStringSlice()
}
