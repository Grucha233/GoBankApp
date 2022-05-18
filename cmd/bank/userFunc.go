package bank

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("123")

type Claims struct {
	Name string `json:"Name"`
	jwt.StandardClaims
}

//This function inserts user data into database
func UserRegister(w http.ResponseWriter, r *http.Request) {
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
func Signin(w http.ResponseWriter, r *http.Request) {
	//Connect to the database
	db, err1 := ConGorm()
	if err1 != nil {
		fmt.Println("Problem with gorm connect")
		w.WriteHeader(http.StatusBadRequest)
		return
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Checking if given password is correct
	if err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form.Get("Password"))); err2 != nil {
		fmt.Println("Wrong credentials, please try again2")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	////////////////////////////////Creating a JWT tocken////////////////////////////////

	//Declare expiration time
	expirationTime := time.Now().Add(1 * time.Hour)

	//Create JWT claims with given UserName and expiry time
	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Creation of JWT tocken based on claims(Username(Name in database) and epiry time)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Converting token to String
	tokenString, err := token.SignedString(jwtKey)
	//If something happens then return error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Sets Cookie with JWT
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	fmt.Println("Succesfully login")
}

//Function checks if JWT is valid and if yup then gives permission to other functionalities of app
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	//Return of the original endpoint if JWT is Valid
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Geting JWT from Cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//Geting JWT into variable
		tokenString := c.Value

		claims := &Claims{}

		//Parse the JWT and store in claims object
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if token.Valid {
			endpoint(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
