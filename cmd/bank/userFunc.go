package bank

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "123"

//This function inserts user data into database
func UserRegister(r *http.Request) {
	//Connect to the database
	db, err1 := ConGorm()
	if err1 != nil {
		fmt.Println("Problem with gorm connect")
	}
	//Parsing POST request
	err2 := r.ParseForm()
	if err2 != nil {
		panic(err2)
	}
	//Checking if user alredy exists in database
	var userCheck Users
	db.First(&userCheck, "Name = ?", r.Form.Get("Name"))
	if userCheck.Name != "" {
		fmt.Println("User alredy exists in database please try again")
		return
	}
	//Hassing password and insert data to the database
	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(r.Form.Get("Password")), 8)
	user := NewUser(r.Form.Get("Name"), r.Form.Get("Currency"), string(hashedPassword))
	db.Create(&user)
}

//This function auth user, and creates JWT for him(1h)
func UserLogin(r *http.Request) (string, error) {
	//Connect to the database
	db, err1 := ConGorm()
	if err1 != nil {
		fmt.Println("Problem with gorm connect")
		return "", err1
	}
	//Parsing POST request
	err2 := r.ParseForm()
	if err2 != nil {
		panic(err2)
	}
	//Checking if Name exists in database
	var user Users
	db.First(&user, "Name = ?", r.Form.Get("Name"))

	if user.ID == 0 {
		fmt.Println("Wrong credentials, please try again")
		return "", err1
	}
	//Checking if given password is correct
	if err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form.Get("Password"))); err2 != nil {
		fmt.Println("Wrong credentials, please try again2")
		return "", err2
	}

	fmt.Println("Succesfully login")
	//Creating a JWT tocken
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		fmt.Println("Something Went Wrong")

	}
	fmt.Println(token)
	return token, nil
}
