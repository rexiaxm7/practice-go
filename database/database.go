package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

func Connect(user string, password string, host string, database string, dialect string) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	db, err := gorm.Open(dialect, connectionString)
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
