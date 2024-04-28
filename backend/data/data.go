package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config2 "backend/config"
)

var (
	DB *gorm.DB
)

// DSN is the Data Source Name for the database connection
const (
	DSN = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type Login struct {
	dbName   string
	username string
	password string
	hostname string
	port     string
}

// ConnectToDB connects to the database and returns a pointer to the database
func ConnectToDB() (*gorm.DB, error) {

	config, err := config2.LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		DSN,
		config.DBUsername,
		config.DBPassword,
		config.DBHostname,
		config.DBPort,
		config.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
