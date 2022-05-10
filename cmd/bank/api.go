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
	ConDB()
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Login page - GET")

	case "POST":
		fmt.Fprintf(w, "Login page - POST")

		//connect to database
		db, err := ConDB()

		//insert of the row
		sqlStatement := `
		INSERT INTO users (name, balance, currency, Password )
		VALUES ('pan1', 65, 'zloty','123')`

		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}

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
