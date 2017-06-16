package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	fmt.Printf("enter main<<<<-----------------------\n")
	sql.op
	db, err := sql.Open("mysql", "videoconference:videoconference@tcp(192.168.20.68:3306)/videoconference?charset=utf8")
	if err != nil {
		fmt.Printf("error connecting: %s \n", err.Error())
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("db.Ping(): err=%s \n", err.Error())
		panic(err)
	}

	//rows, err := db.Query("select id, username from user where id = ?", 1)
	rows, err := db.Query("select a.meetingId,a.hostId from meeting a, competence b where a.meetingId = b.meetingId", 1)
	if err != nil {
		//fmt.Printf("db.Query: err=%s \n", err.Error())
		//panic(err)
		log.Fatalf("db.Query error: %s\n", err)
	}

	defer rows.Close()

	var meetingId int
	var hostid int
	for rows.Next() {
		err := rows.Scan(&meetingId, &hostid)
		if err != nil {
			fmt.Printf("rows.Scan: err=%s \n", err.Error())
			panic(err)
		}
		fmt.Printf("meetingId:%d, hostid:%d \n", meetingId, hostid)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf("rows.Err(): err=%s \n", err.Error())
		panic(err)
	}

	fmt.Printf("extit main ----------------------->>>>>\n")
}
