package account

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/html")

	tmp.Execute(w, nil)
}
