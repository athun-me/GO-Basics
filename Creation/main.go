package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Trying to connecte...")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test1")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	Create, err := db.Query("CREATE TABLE table_test2(id INT, name varchar(20), age varchar(10), PRIMARY KEY(id))")

	if err != nil {
		panic(err.Error())
	}

	defer Create.Close()

	fmt.Println("Succcesfully Created")
}
