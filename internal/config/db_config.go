package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(verbose bool, logger bool) *gorm.DB {
	if logger {
		println("[+] Processing: Initializing Database [+]")
	}

	return dbConn(logger)
}

func dbConn(logger bool) *gorm.DB {
	dbHost := viper.GetString("database.host")
	dbName := viper.GetString("database.name")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbPort := viper.GetInt("database.port")

	/* build dsn with the information needed to connect to the database, including the username,
	password, host, port, and database name */
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// establish a connection to the database using the dsn
	db, err := gorm.Open(
		mysql.Open(dsn), &gorm.Config{},
	)
	if err != nil {
		if logger {
			log.Println("[-] Error: error while connecting to the database [-]")
		}

		log.Fatal(err.Error())
	}

	// check if the database is accessible
	_, err = db.DB()
	if err != nil {
		log.Println("[-] Error: Unable to get database instance [-]")
	}

	return db
}
