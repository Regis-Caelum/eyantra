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
	"github.com/gorilla/sessions"
)

var tpl *template.Template

var static_name string

// type table struct {
// 	rows [90]string
// }

//--------------------------------------------------session variables------------------------------------------------------------
var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

//--------------------------------------------------Template initializer---------------------------------------------------------

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

//---------------------------------------------------index handler---------------------------------------------------------------

func index(w http.ResponseWriter, r *http.Request) {
	//creating session
	session, _ := store.Get(r, "login")
	session.Values["authenticated"] = true

	//Saving session value
	session.Save(r, w)

	//Identifying error in parsing the request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//variables for authentication
	var husr, hpass, email, pass string

	// checking if user wants the log in for root dashboard or hospital dashboard
	if r.FormValue("button") == "hospital" {
		husr = r.PostFormValue("huser")
		static_name = husr
		hpass = r.PostFormValue("hpass")

		check := HospitalLoginCheck(husr, hpass)

		if check {
			fmt.Println(check)
			http.Redirect(w, r, "/hospital", http.StatusFound)
		} else {
			fmt.Println(check)
			fmt.Fprintf(w, "USER NOT FOUND\n")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	if r.FormValue("buttons") == "root" {
		email = r.PostFormValue("auser")
		pass = r.PostFormValue("apass")
		chck := AdminLoginCheck(string(email), string(pass))

		if chck {
			http.Redirect(w, r, "/main", http.StatusFound)
		} else {
			fmt.Println(chck)
			fmt.Fprintf(w, "USER NOT FOUND\n")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

}

// MAIN FUNCTION
func main() {
	//Indicating that the server is up and running
	fmt.Println("Serving at localhost:8080...")

	//Saving templates directory in the fileServer variable
	fileServer := http.FileServer(http.Dir("./templates"))

	//Adding routes
	http.Handle("/", fileServer)
	http.HandleFunc("/login", index)
	http.HandleFunc("/main", newentry)
	http.HandleFunc("/done", temp)
	http.HandleFunc("/hospital", hospital)
	http.HandleFunc("/hospitaldone", hospitaldone)
	// http.HandleFunc("/view", viewhandler)
	// http.HandleFunc("/error", search)

	//Listening on port 8080
	http.ListenAndServe(":8080", nil)
}

// func search(w http.ResponseWriter, r *http.Request) {
//  	//Getting cookie
// 	session, _ := store.Get(r, "login")
// 	//Checking for authentication
// 	if auth, err := session.Values["authenticated"].(bool); !err || !auth {
// 		http.Error(w, "Forbidden", http.StatusForbidden)
// 		return
// 	}
// 	//If condition satisfies executing the template
// 	tpl.ExecuteTemplate(w, "search.html", nil)
// }

// Check in database for email and it's corresponding password
func AdminLoginCheck(mail string, pass string) bool {
	//variable for return
	status := false

	//opening connection to mysql database
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	//checking for error in the connection
	if err != nil {
		fmt.Printf("not connected")
	}
	//delayed the closing connection for database
	defer db.Close()

	//SQL query
	query := "SELECT Passwords FROM test_schema.database WHERE Email ='" + mail + "'"

	//Executing query
	res, err := db.Query(query)

	//Checking for errors
	if err != nil {
		panic(err.Error())
	}

	//Going the thr result set rows and storing every row in the string variable
	for res.Next() {
		var password string
		//storing row in the string variable and checking for errors in the process simultaneously
		if err := res.Scan(&password); err != nil {
			log.Fatal(err)
		}

		//If the password matches with the entered value then changing the state of status variable
		if password == pass {
			status = true
		}
	}

	//Checking for error in the res variable
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return status
}

// New Entry
func newentry(w http.ResponseWriter, r *http.Request) {

	//Getting cookie
	session, _ := store.Get(r, "login")
	//Checking for authentication
	if auth, err := session.Values["authenticated"].(bool); !err || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//If condition satisfies executing the template
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func temp(w http.ResponseWriter, r *http.Request) {
	//Variables for storing the data in the database
	district := r.FormValue("district")
	pin := r.FormValue("pin")
	hospital_name := r.FormValue("name")

	//Entering the new entry and checking for error simultaneously by passing a flag
	chck := newentity(district, pin, hospital_name)

	//If flag holds redirecting to rew route
	if chck {
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
	if errs != nil {
		fmt.Println(errs)
	}
	return stat
}

//view table
// func viewhandler(w http.ResponseWriter, r *http.Request) {
//	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")
//
// 	if err != nil {
// 		fmt.Printf("not connected")
// 	}
// 	defer db.Close()
//	query := "SELECT namess FROM test_schema.data WHERE pincode='" + r.FormValue("pin") + "';"
//	res1,_ := db.Query(query)
//	var named string
//  for res1.Next() {
//
// 	 if err := res1.Scan(&named); err != nil {
// 		 log.Fatal(err)
// 	 }
//
//  }
// 	res := view(named)
//  i := 0
// 	for res.Next() {
// 		var row string
// 		if err := res.Scan(&row); err != nil {
// 			log.Fatal(err)
// 		}
// 		p = table{row[i]: row}
// 		i = i+1
// 		fmt.Fprintf(w, row+"\n")
// 	}
// }

// func view(name string) *sql.Rows {
// 	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

// 	if err != nil {
// 		fmt.Printf("not connected")
// 	}
// 	defer db.Close()
// 	query := "SELECT namess, oxygen_beds, ventilator_beds, normal_bed FROM test_schema.data WHERE namess ='" + name + "';"

// 	res, _ := db.Query(query)

// 	return res
// }

//Hospital handler

func hospital(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	if auth, err := session.Values["authenticated"].(bool); !err || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "hospital.html", nil)
}

func hospitaldone(w http.ResponseWriter, r *http.Request) {
	oxygen := r.FormValue("oxy")
	vent := r.FormValue("vent")
	norm := r.FormValue("norm")
	chck := addbeds(oxygen, vent, norm, static_name)

	if chck {
		http.Redirect(w, r, "/hospital", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusForbidden)
	}
}

func addbeds(oxy string, vent string, norm string, name string) bool {
	stat := true
	//fmt.Printf(" %v, %v, %v\n", oxy, vent, norm)
	query := "UPDATE test_schema.hospital SET oxygen_beds='" + oxy + "', ventilator_beds='" + vent + "', normal_bed='" + norm + "' WHERE namess='" + name + "';"
	//fmt.Println(query)
	//query := "UPDATE test_schema.hospital SET oxygen_beds='0'"
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

func HospitalLoginCheck(usr string, hpass string) bool {
	stat := false
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/test_schema")

	if err != nil {
		fmt.Printf("not connected")
	}
	defer db.Close()
	query := "SELECT passwords FROM test_schema.hospital WHERE namess ='" + usr + "'"
	res, _ := db.Query(query)

	for res.Next() {
		var passe string
		if err := res.Scan(&passe); err != nil {
			log.Fatal(err)
		}
		fmt.Println(passe)
		if passe == hpass {
			stat = true
		}
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return stat
}

//---------------------------------------------------------Random Sting Generator------------------------------------------------------------------------
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
