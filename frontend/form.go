// forms.go
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// caches templates
var templates = template.Must(template.ParseFiles("forms.html"))

type ContactDetails struct {
	firstName string
	lastName  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, v any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func submitForm(w http.ResponseWriter, r *http.Request) {

	// CHECK IF METHOD IS POST
	if r.Method != http.MethodPost {
		templates.ExecuteTemplate(w, "forms.html", nil)
		return
	}

	// EXTRACT SUBMISSION
	details := ContactDetails{
		firstName: r.FormValue("first_name"),
		lastName:  r.FormValue("last_name"),
	}

	// CONVERT SUBMISSION TO JSON
	req := fmt.Sprintf(`{"first_name": "%s", "last_name": "%s"}`, details.firstName, details.lastName)
	jsonBody := []byte(req)
	bodyReader := bytes.NewReader(jsonBody)

	// SUBMIT TO BACKEND API
	requestURL := fmt.Sprintf("http://localhost:%d/api/v1/create", 3000)
	resp, err := http.Post(requestURL, "applicaton/json", bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	// REFRESH TEMPLATE
	renderTemplate(w, "forms", struct{ Success bool }{true})
}

func main() {

	http.HandleFunc("/register", submitForm)

	portNumber := ":8080"
	log.Printf("Web server running on port %s", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
