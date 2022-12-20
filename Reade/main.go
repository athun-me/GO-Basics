package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Trying to connect....")

	
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test1")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insertion, err := db.Query("INSERT INTO users VALUES('aTHUNL lal')")

	if err != nil {
		panic(err.Error())
	}

	defer insertion.Close()

	fmt.Println("Succesfully Connected into users ")
}
