package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name      string
	age       uint16
	money     int16
	avgGrades float64
	happines  float64
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go go go")
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8000", nil)
}

func main() {
	handleRequest()
}
