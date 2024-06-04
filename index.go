package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"	

)
var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	
	db, err = sql.Open("sqlite3", "./database.db") // => nom de database 
	if err != nil {
		log.Fatal(err)
	}

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
 func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/newpost", newPost)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}