package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	dbpassword := "root"
	dbhost := "127.0.0.1"

	db, err := gorm.Open("mysql", "root:"+dbpassword+"@tcp("+dbhost+":3306)/hsensedb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(structs.Person{}, structs.Administrator{}, structs.GPS{}, structs.GPSHistory{}, structs.Message{}, structs.Whatsapp{}, structs.Telegram{}, structs.Contact{}, structs.Picture{}, structs.Keylog{})
	return db
}
