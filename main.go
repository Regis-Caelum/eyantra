package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// fmt.Fprintf(w, "POST request successful\n")
	email := r.FormValue("user")
	pass := r.FormValue("pass")

	chck := check(string(email), string(pass))

	if chck {
		tpl.ExecuteTemplate(w, "main.html", nil)
		//http.ServeFile(w, r, "./templates/main.html")
	} else {
		fmt.Fprintf(w, "USER NOT FOUND\n")
	}

}

func main() {
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)
	http.HandleFunc("/main", formHandler)
	http.ListenAndServe(":8080", nil)
}

// Check in database for email and it's corresponding password
func check(mail string, pass string) bool {
	status := false
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}

	defer db.Close()

	query := "SELECT Passwords FROM test_schema.database WHERE Email ='" + mail + "'"
	res, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	for res.Next() {
		var name string
		if err := res.Scan(&name); err != nil {
			log.Fatal(err)
		}
		if name == pass {
			status = true
		}
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return status
}
