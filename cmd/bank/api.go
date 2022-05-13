package bank

import (
	"fmt"
	"log"
	"net/http"
)

//
// Main api function, handles all endpoints
func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

//
// Allows to register, "/register"
func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Login page - GET")

	case "POST":
		fmt.Fprintf(w, "Login page - POST")
		UserRegister(r)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

//
// Allows users to log in to the bank, "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login page")
	fmt.Println("Endpoint Hit: Login")
}

//
// Homepage handling function, "/"
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// usrStr := `{
// 	"Name":"adas",
// 	"Balance": "170",
// 	"Currency": "zloty",
// 	"Password": "123"
//  }`
