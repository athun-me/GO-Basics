package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Trying to connect")
	db, err := sql.Open("mysql", "root:@tcp(localhost)/test1")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO table_test2 VALUES('2', 'Arjun', '19')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Succesfully Inserted")
}
