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
// handling "/adas" endpoint function
func AdasFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to Adaś!!!")

}

//
// Homepage handling function
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
