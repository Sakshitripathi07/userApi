package database

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"userApi/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//NewDB creates a database and connects to the database.
func NewDB() {
	host := "localhost"
	port := 5433
	user := "apiuser"
	password := "root"
	dbname := "apidb"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	DB = db
}
