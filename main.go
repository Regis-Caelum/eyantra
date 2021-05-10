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

// MAIN FUNCTION
func main() {
	fmt.Println("Serving at localhost:8080...")
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)
	http.HandleFunc("/main", formHandler)
	http.HandleFunc("/new", newentry)
	http.HandleFunc("hospital", hospital)
	http.HandleFunc("/view", viewhandler)
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

// New Entry
func newentry(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "main.html", nil)

	district := r.FormValue("district")
	pin := r.FormValue("pin")
	hospital_name := r.FormValue("name")

	chck := newentity(district, pin, hospital_name)

	if chck {
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		fmt.Fprintf(w, "Can't add the entry something occured\n")
	}
}

// Connection to New database on same server
func newentity(district string, pin string, hospital_name string) bool {
	stat := true
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}

	defer db.Close()

	query := "INSERT INTO test_schema.data (namess, pincode, district) SELECT * FROM (SELECT '" + hospital_name + "' AS namess, '" + pin + "' AS pincode, '" + district + "' AS district) AS temp WHERE NOT EXISTS( SELECT namess FROM test_schema.data WHERE district='" + district + "' AND pincode='" + pin + "');"

	db.Query("INSERT INTO test_schema.hospital (namess) SELECT * FROM (SELECT '" + hospital_name + "' AS namess ) AS temp WHERE NOT EXISTS( SELECT namess FROM test_schema.data;")
	_, err = db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		stat = false
	}
	return stat
}

//view table
func viewhandler(w http.ResponseWriter, r *http.Request) {
	res := view(r.FormValue("pin"))

	for res.Next() {
		var table string
		if err := res.Scan(&table); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, table+"\n")
	}
}

func view(pin string) *sql.Rows {
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}
	defer db.Close()
	query := "SELECT name FROM test_schema.data WHERE pincode ='" + pin + "';"

	res, _ := db.Query(query)

	return res
}

//Hospital handler
func hospital(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hospital.html", nil)

	oxygen := r.FormValue("oxy")
	vent := r.FormValue("vent")
	norm := r.FormValue("norm")

	chck := addbeds(oxygen, vent, norm, " ")

	if chck {
		fmt.Fprintf(w, "Entered Sucessfully\n")
	}
}

func addbeds(oxy string, vent string, norm string, name string) bool {
	stat := true
	query := "UPDATE test_schema.hospital SET oxygen_beds='" + oxy + "', ventilator_beds='" + vent + "', normal_beds='" + norm + "' WHERE namess='" + name + "'"
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}
	defer db.Close()

	db.Query(query)
	return stat
}
