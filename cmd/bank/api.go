package bank

import (
	"fmt"
	"log"
	"net/http"
)

//
// Main api function, handles all endpoints
func HandleRequests() {
	http.Handle("/", IsAuthorized(HomePage))
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/check", Check)
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
		UserRegister(w, r)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

//
// Allows users to log in to the bank, "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Welcome to Login page, please pass credentials to LogIn")

	case "POST":
		fmt.Println("Trying to log in")
		Signin(w, r)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Checking function")
}

//
// Homepage handling function, "/"
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
