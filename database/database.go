package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Connect(user string, password string, host string, database string, dialect string) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}
