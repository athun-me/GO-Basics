package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
}

func main() {
	fmt.Println("Trying to connect....")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/company")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	Read, err := db.Query("SELECT first_name, last_name, email, gender from test1")

	if err != nil {
		panic(err.Error())
	}

	for Read.Next() {
		var user Users

		err = Read.Scan(&user.First_name, &user.Last_name, &user.Email, &user.Gender)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.First_name, user.Last_name, user.Email, user.Gender)
	}

	fmt.Println("Succesfully  Read Users ")
}
