package main

import (
	"ORMTest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UserInfo struct {
	gorm.Model
	Name string
	Age  int
}
type AppUsing struct {
	Id          int
	AppName     string
	SourceName  string
	SourceValue string
}

func main() {
	db, err := models.MysqlConnect()
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	if db.HasTable(&AppUsing{}) {
		db.Create(&AppUsing{AppName: "Hello", SourceName: "foo", SourceValue: "bar"})
	} else {
		db.CreateTable(&AppUsing{})
	}

}
