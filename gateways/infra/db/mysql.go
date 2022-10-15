package db

import (
	"hex-grpc/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// user := os.Getenv("DB_USERNAME")
	// pass := os.Getenv("DB_PASSWORD")
	// protocol := os.Getenv("DB_PROTOCOL")
	// dbname := os.Getenv("DB_DATABASE")

	// dsn := user + ":" + pass + "@" + protocol + "/" + dbname + "?charset=utf8&parseTime=true"
	// connection, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:               dsn,
	// 	DefaultStringSize: 256,
	// }), &gorm.Config{})

	connection, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	err = connection.AutoMigrate(&entities.Article{})
	if err != nil {
		panic("failed to migrate")
	}

	return connection
}
