package bank

import (
	"database/sql"
	"fmt"
)

//
//information needet to connect to database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "bank_database"
)

//
//Connects to the database
func ConDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

type User struct {
	id         int
	full_name  string
	balance    float64
	currency   string
	created_at string
}

type Entries struct {
	id         int
	account_id int
	ammount    float64
	created_at string
}

type Transfers struct {
	id              int
	from_accuont_id int
	to_account_id   int
	ammount         float64
	created_at      string
}
