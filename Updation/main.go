package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Trying to connect")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test1")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	update, err := db.Query("UPDATE table_test2 SET age='23' WHERE id='1'")

	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
	fmt.Println("Successfully updated")
}
