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
	http.HandleFunc("/adas", AdasFunc)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

//
// prints on AdasPage Welcome
func AdasFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to Adaś!!!")

}

//
// homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
