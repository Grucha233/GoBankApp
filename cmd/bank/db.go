package bank

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
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

//Creates pointer to the database (connects to the database)
func ConDB() (sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return *db, err
}

type Users struct {
	gorm.Model
	ID         int64
	Name       string
	Balance    sql.NullFloat64
	Currency   string
	Created_at string
	Password   string
}

type Transfers struct {
	gorm.Model
	ID              int64
	From_account_id int64
	To_account_id   int64
	Ammount         float64
	Created_at      string
}

type Entries struct {
	gorm.Model
	ID         int64
	Account_id int64
	Ammount    float64
	Created_at string
}
