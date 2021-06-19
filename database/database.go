package database

import "github.com/jinzhu/gorm"

var Database *gorm.DB

func Connect() *gorm.DB {
	CONNECT := "test:test@tcp(localhost:3306)/practice"
	db, err := gorm.Open("mysql", CONNECT)
	Database = db
	if err != nil {
		panic(err.Error())
	}

	return db
}

func Close() {
	err := Database.Close()
	if err != nil {
		panic(err.Error())
	}
}
