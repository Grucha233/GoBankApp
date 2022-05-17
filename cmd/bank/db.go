package bank

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Creates pointer to the database gorm (connects to the database)
func ConGorm() (gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=bank_database2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return *db, err
}

//Migrates models to the database (should be userd when program starts)
func AutoMigrate() {
	db, err := ConGorm()
	if err != nil {
		fmt.Println("Connection problem")
	}
	var user Users
	var entrie Entries
	var transfer Transfers
	db.AutoMigrate(&user, &entrie, &transfer)
}

type Users struct {
	gorm.Model
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Balance  float64
	Currency string
	Password string
}

//Constructor of new user
func NewUser(name string, currency string, password string) Users {
	nu := Users{
		Name:     name,
		Currency: currency,
		Balance:  0,
		Password: password,
	}
	return nu
}

type Transfers struct {
	gorm.Model
	ID              int64 `gorm:"primaryKey"`
	From_account_id int64
	To_account_id   int64
	Ammount         float64
}

type Entries struct {
	gorm.Model
	ID         int64 `gorm:"primaryKey"`
	Account_id int64
	Ammount    float64
}

func NewUserFromJson(jsonStr string) (Users, error) {
	var newUser Users
	err := json.Unmarshal([]byte(jsonStr), &newUser)
	if err != nil {
		err = fmt.Errorf("unable to unmarshal JSON", err.Error())
		return Users{}, err
	}
	return newUser, nil
}
