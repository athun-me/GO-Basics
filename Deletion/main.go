package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Trying to connect..")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test1")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	Delete, err := db.Query("DELETE FROM users WHERE name='ELLIOT'")

	if err != nil {
		panic(err.Error())
	}

	defer Delete.Close()

	fmt.Println("Succesfully deleted")
}
