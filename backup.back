package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./templates/index.html")
	// tpl.ExecuteTemplate(w, "index.html", nil)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	var husr, hpass, email, pass string

	fmt.Println(r.PostFormValue("button"))

	// fmt.Fprintf(w, "POST request successful\n")
	if r.FormValue("button") == "hospital" {
		husr = r.PostFormValue("huser")
		hpass = r.PostFormValue("hpass")

		check := hcheck(husr, hpass)
		fmt.Println("Reached checkpoint")
		fmt.Println(check)
		if check {
			fmt.Println(check)
			http.Redirect(w, r, "/hospital", http.StatusFound)
			//http.ServeFile(w, r, "./templates/main.html")
		} else {
			fmt.Println(check)
			fmt.Fprintf(w, "USER NOT FOUND\n")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	if r.FormValue("buttons") == "root" {
		email = r.PostFormValue("auser")
		pass = r.PostFormValue("apass")
		chck := check(string(email), string(pass))

		if chck {
			fmt.Println(chck)
			http.Redirect(w, r, "/main", http.StatusFound)
			//http.ServeFile(w, r, "./templates/main.html")
		} else {
			fmt.Println(chck)
			fmt.Fprintf(w, "USER NOT FOUND\n")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

}

// MAIN FUNCTION
func main() {
	fmt.Println("Serving at localhost:8080...")
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)
	//http.HandleFunc("/", index)
	http.HandleFunc("/login", index)
	http.HandleFunc("/main", newentry)
	http.HandleFunc("/done", temp)
	http.HandleFunc("/hospital", hospital)
	http.HandleFunc("/hospitaldone", hospitaldone)
	http.HandleFunc("/view", viewhandler)
	http.HandleFunc("/error", errored)
	http.ListenAndServe(":8080", nil)
}

func errored(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Error")
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
		var passe string
		if err := res.Scan(&passe); err != nil {
			log.Fatal(err)
		}
		if passe == pass {
			status = true
			fmt.Println(passe)
		}
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return status
}

// New Entry
func newentry(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./templates/main.html")
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func temp(w http.ResponseWriter, r *http.Request) {

	district := r.FormValue("district")
	pin := r.FormValue("pin")
	hospital_name := r.FormValue("name")

	//fmt.Printf("%v %v %v\n", district, pin, hospital_name)
	chck := newentity(district, pin, hospital_name)

	if chck {
		//fmt.Printf("working...\n")
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
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

	_, errs := db.Query(" INSERT INTO test_schema.hospital (namess, passwords, oxygen_beds, ventilator_beds, normal_bed) SELECT * FROM (SELECT '" + hospital_name + "' AS namess, '" + RandStringRunes(4) + "' AS passwords, '0' AS oxygen_beds,'0' AS ventilator_beds,'0' AS normal_bed) AS tmp WHERE NOT EXISTS ( SELECT namess FROM test_schema.hospital WHERE namess = '" + hospital_name + "' );")
	_, err = db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		stat = false
	}
	if err != nil {
		fmt.Println(errs)
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
}

func hospitaldone(w http.ResponseWriter, r *http.Request) {
	oxygen := r.FormValue("oxy")
	vent := r.FormValue("vent")
	norm := r.FormValue("norm")
	//fmt.Printf("%T", oxygen)
	chck := addbeds(oxygen, vent, norm, " ")

	if chck {
		http.Redirect(w, r, "/hospital", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func addbeds(oxy string, vent string, norm string, name string) bool {
	stat := true
	query := "UPDATE test_schema.hospital SET oxygen_beds='" + oxy + "', ventilator_beds='" + vent + "', normal_bed='" + norm + "' WHERE namess='" + name + "'"
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}
	defer db.Close()

	_, err = db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		stat = false
	}

	return stat
}

func hcheck(usr string, hpass string) bool {
	stat := false
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}
	defer db.Close()
	query := "SELECT passwords FROM test_schema.hospital WHERE namess ='" + usr + "'"
	res, _ := db.Query(query)

	fmt.Println("Reached checkpoint")
	for res.Next() {
		var passe string
		if err := res.Scan(&passe); err != nil {
			log.Fatal(err)
		}
		fmt.Println(passe)
		if passe == hpass {
			stat = true
			fmt.Println(passe)
		}
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return stat
}

//-------------------------------------------------------------------------------------------------------------------------------------

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
