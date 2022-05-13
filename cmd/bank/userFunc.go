package bank

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//Inserts user data into database
func UserRegister(r *http.Request) {
	db, err := ConGorm()
	if err != nil {
		fmt.Println("Problem with gorm connect")
	}
	u := GetUserData(r)
	db.Create(&u)
}

func GetUserData(r *http.Request) Users {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("Password")), 8)
	user := NewUser(r.Form.Get("Name"), r.Form.Get("Currency"), string(hashedPassword))
	return user
}
