package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "name:password/database?charset=uft8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to the database ...")
	db = database

}

func GetDB() *gorm.DB {
	return db
}
