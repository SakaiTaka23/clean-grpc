package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	dbname := os.Getenv("DB_DATABASE")

	dsn := user + ":" + pass + "@" + protocol + "/" + dbname + "?charset=utf8&parseTime=true"
	connection, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	return connection
}
