package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var tpl *template.Template

var static_name string

type hospital_list struct {
	district, pincode, name string
}

func (p *hospital_list) District() string { return p.district }
func (p *hospital_list) Pincode() string  { return p.pincode }
func (p *hospital_list) Name() string     { return p.name }

//--------------------------------------------------session variables------------------------------------------------------------
var (
	key    = []byte("secret-key")
	store  = sessions.NewCookieStore(key)
	emails string
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
		emails = string(husr)
		//fmt.Println(emails, husr)
		check := HospitalLoginCheck(husr, hpass)

		if check {
			http.Redirect(w, r, "/hospital", http.StatusFound)
		} else {
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
			fmt.Fprintf(w, "USER NOT FOUND\n")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

}

// MAIN FUNCTION
func main() {
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// 	panic(err)
	// }

	// //Indicating that the server is up and running
	// fmt.Printf("Serving at http://localhost:%v\n", listener.Addr().(*net.TCPAddr).Port)
	// c := exec.Command("python3", "./webscrapper/web-s.py")

	// if err := c.Run(); err != nil {
	// 	fmt.Println("Error: ", err.Error())
	// }
	//readData()
	fmt.Println("Updated to latest records...")
	//Saving templates directory in the fileServer variable
	fileServer := http.FileServer(http.Dir("./templates"))

	//Adding routes
	http.Handle("/", fileServer)
	http.HandleFunc("/login", index)
	http.HandleFunc("/main", newentry)
	http.HandleFunc("/done", temp)
	http.HandleFunc("/searching", searching)
	http.HandleFunc("/search", search)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/hospital", hospital)
	http.HandleFunc("/hospitaldone", hospitaldone)

	//Listening on any free port given by the OS
	// http.Serve(listener, nil)
	port := os.Getenv("PORT")
	fmt.Printf("Starting server at port available %v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	session.Values["authenticated"] = false

	//Saving session value
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

type Info struct {
	name string
	oxy  string
	norm string
	vent string
}

func (p *Info) Name() string { return p.name }
func (p *Info) Oxy() string  { return p.oxy }
func (p *Info) Norm() string { return p.norm }
func (p *Info) Vent() string { return p.vent }

var info, info2 []Info

func searching(w http.ResponseWriter, r *http.Request) {
	//If condition satisfies executing the template
	tpl.ExecuteTemplate(w, "search.html", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	//opening connection to mysql database
	db, err := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

	//checking for error in the connection
	if err != nil {
		fmt.Printf("not connected")
	}
	//delayed the closing connection for database
	defer db.Close()

	text := r.FormValue("text")

	if text == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	q := "SELECT namess FROM sql4413785.data WHERE pincode='" + text + "';"

	result, _ := db.Query(q)

	var list []string
	for result.Next() {
		var temp string
		err := result.Scan(&temp)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, temp)
	}

	for _, s := range list {
		query := "SELECT * FROM sql4413785.hospital WHERE namess='" + s + "';"

		res, _ := db.Query(query)

		for res.Next() {
			var nam, o, v, no, temp string
			err = res.Scan(&nam, &temp, &o, &v, &no)
			if err != nil {
				panic(err.Error())
			}
			details := Info{
				name: nam,
				oxy:  o,
				vent: v,
				norm: no,
			}

			info = append(info, details)
		}
	}

	tpl.ExecuteTemplate(w, "search.html", &info)
	info = nil

}

// Check in database for email and it's corresponding password
func AdminLoginCheck(mail string, pass string) bool {
	//variable for return
	status := false

	//opening connection to mysql database
	db, err := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

	//checking for error in the connection
	if err != nil {
		fmt.Printf("not connected")
	}
	//delayed the closing connection for database
	defer db.Close()

	//SQL query
	query := "SELECT Passwords FROM sql4413785.database WHERE Email ='" + mail + "'"

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

var hospital_lists []hospital_list

// New Entry
func newentry(w http.ResponseWriter, r *http.Request) {

	//Getting cookie
	session, _ := store.Get(r, "login")
	//Checking for authentication
	if auth, err := session.Values["authenticated"].(bool); !err || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	db, _ := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

	defer db.Close()

	query := "SELECT * FROM sql4413785.data;"

	res, _ := db.Query(query)

	i := 0
	for res.Next() {
		var dist, pin, nam string
		err := res.Scan(&dist, &pin, &nam)
		hospital_lists = append(hospital_lists, hospital_list{
			district: dist,
			pincode:  pin,
			name:     nam,
		})
		i = i + 1
		if err != nil {
			log.Println(err.Error())
			return
		}

	}

	//If condition satisfies executing the template
	tpl.ExecuteTemplate(w, "main.html", &hospital_lists)
	hospital_lists = nil
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
	db, err := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

	if err != nil {
		fmt.Printf("not connected")
	}

	defer db.Close()

	query := "INSERT INTO sql4413785.data (namess, pincode, district) SELECT * FROM (SELECT '" + hospital_name + "' AS namess, '" + pin + "' AS pincode, '" + district + "' AS district) AS temp WHERE NOT EXISTS( SELECT namess FROM sql4413785.data WHERE district='" + district + "' AND pincode='" + pin + "');"

	password := RandStringRunes(4)

	_, errs := db.Query(" INSERT INTO sql4413785.hospital (namess, passwords, oxygen_beds, ventilator_beds, normal_bed) SELECT * FROM (SELECT '" + hospital_name + "' AS namess, '" + password + "' AS passwords, '0' AS oxygen_beds,'0' AS ventilator_beds,'0' AS normal_bed) AS tmp WHERE NOT EXISTS ( SELECT namess FROM sql4413785.hospital WHERE namess = '" + hospital_name + "' );")
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

//Hospital handler
func hospital(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	if auth, err := session.Values["authenticated"].(bool); !err || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	db, _ := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

	query := "SELECT * FROM sql4413785.hospital WHERE namess='" + emails + "';"
	//fmt.Println(query)
	res, _ := db.Query(query)

	for res.Next() {
		var nam, o, v, no, temp string
		err := res.Scan(&nam, &temp, &o, &v, &no)
		if err != nil {
			panic(err.Error())
		}
		details := Info{
			name: nam,
			oxy:  o,
			vent: v,
			norm: no,
		}

		info2 = append(info2, details)
	}

	tpl.ExecuteTemplate(w, "hospital.html", &info2)
	info2 = nil
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
	var query string
	if oxy == "" && vent == "" && norm == "" {
		query = " "
	} else if oxy == "" && vent == "" {
		query = "UPDATE sql4413785.hospital SET normal_bed='" + norm + "' WHERE namess='" + name + "';"
	} else if oxy == "" && norm == "" {
		query = "UPDATE sql4413785.hospital SET ventilator_beds='" + vent + "' WHERE namess='" + name + "';"
	} else if vent == "" && norm == "" {
		query = "UPDATE sql4413785.hospital SET oxygen_beds='" + oxy + "' WHERE namess='" + name + "';"
	} else if oxy == "" {
		query = "UPDATE sql4413785.hospital SET ventilator_beds='" + vent + "', normal_bed='" + norm + "' WHERE namess='" + name + "';"
	} else if vent == "" {
		query = "UPDATE sql4413785.hospital SET oxygen_beds='" + oxy + "', normal_bed='" + norm + "' WHERE namess='" + name + "';"
	} else if norm == "" {
		query = "UPDATE sql4413785.hospital SET oxygen_beds='" + oxy + "', ventilator_beds='" + vent + "', normal_bed='" + norm + "' WHERE namess='" + name + "';"
	} else {
		query = "UPDATE sql4413785.hospital SET oxygen_beds='" + oxy + "', ventilator_beds='" + vent + "' WHERE namess='" + name + "';"
	}

	db, err := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

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
	db, err := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

	if err != nil {
		fmt.Printf("not connected")
	}
	defer db.Close()
	query := "SELECT passwords FROM sql4413785.hospital WHERE namess ='" + usr + "'"
	res, _ := db.Query(query)

	for res.Next() {
		var passe string
		if err := res.Scan(&passe); err != nil {
			log.Fatal(err)
		}

		if string(passe) == hpass {
			stat = true
		}
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return stat
}

//---------------------------------------------------------Random Sting Generator------------------------------------------------------------------------

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//--------------------------------------------------------Data filler-----------------------------------------------------------------------
// type data struct {
// 	sr   string
// 	name string
// 	o2   string
// 	non2 string
// 	icu  string
// 	vent string
// }

// var temps []data

// func readData() {

// 	db, err := sql.Open("mysql", "sql4413785:xp7qaQRnNs@tcp(sql4.freesqldatabase.com)/sql4413785")

// 	defer db.Close()

// 	csvFile, err := os.Open("./webscrapper/data/test.csv")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Successfully Opened CSV file")
// 	defer csvFile.Close()

// 	csvLines, err := csv.NewReader(csvFile).ReadAll()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, line := range csvLines {
// 		//fmt.Println(line)
// 		emp := data{
// 			sr:   line[0],
// 			name: line[1],
// 			o2:   line[2],
// 			non2: line[3],
// 			icu:  line[4],
// 			vent: line[5],
// 		}
// 		temps = append(temps, emp)
// 	}
// 	for _, x := range temps {
// 		//fmt.Println(x)
// 		query := "UPDATE sql4413785.hospital SET oxygen_beds='" + x.o2 + "', ventilator_beds='" + x.vent + "', normal_bed='" + x.non2 + "' WHERE namess='" + x.name + "';"
// 		//fmt.Println(query)
// 		_, err := db.Query(query)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 	}
// }
