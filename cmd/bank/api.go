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

	fmt.Fprintf(w, "Register page")
	fmt.Println("Endpoint Hit: Register")

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
