package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	type data struct {
		Name string
	}

	err = t.ExecuteTemplate(w, fileName, data{"Please log in here "})
	if err != nil {
		fmt.Println(err)
		return
	}

}

var userDB = map[string]string{
	"Athunlal": "goodPassword",
	"Arjun":    "arjun123",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	passwrod := r.FormValue("password")
	// fmt.Println(username, passwrod)

	if userDB[username] == passwrod {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You've now logged in. Welcome to the main page..")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Didn't find any matched credensials...!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Somthing went wrong.....")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
