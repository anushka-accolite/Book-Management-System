package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Pass@accolite$3@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//db.Exec("SET time_zone = '+00:00'")

	db = d
}

func GetDB() *gorm.DB {
	return db
}
