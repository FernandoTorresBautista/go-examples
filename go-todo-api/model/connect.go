package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connedted to the db")
	con = db
	return db
}

//
