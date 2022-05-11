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
		db, err := ConGorm()
		if err != nil {
			fmt.Println("Problem with gorm connect")
		}
		result := map[string]interface{}{}
		db.Model(&Users{}).First(&result)
		fmt.Println(result)

	case "POST":
		fmt.Fprintf(w, "Login page - POST")
		db, err := ConGorm()
		if err != nil {
			fmt.Println("Problem with gorm connect")
		}
		userStr := `{
			"Name":"Adas2",
			"Currency": "USD",
			"Password": "12345"
		 }`
		nu, err := NewUserFromJson(userStr)
		if err != nil {
			fmt.Println("Problem with UserString")
		}
		db.Create(&nu)

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
