package main

import (
	"log"

	"api-project/internal/config"
	"api-project/internal/domain/models"

	"gorm.io/gorm"
)

func main() {
	// Configure Viper settings for reading config file
	config.ConfigureViper()

	// Initialize the database connection
	db := config.InitDB(true, false)

	// Perform database model migration
	MakeMigrations(db)

}
func MakeMigrations(db *gorm.DB) {

	println("[+][+] Processing: Migrating User Model [+][+]")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		println("[-][-] Failed: User migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating UserBalance Model [+][+]")
	if err := db.AutoMigrate(&models.UserBalance{}); err != nil {
		println("[-][-] Failed: UserBalance migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TransactionLog Model [+][+]")
	if err := db.AutoMigrate(&models.TransactionLog{}); err != nil {
		println("[-][-] Failed:  TransactionLog migration failed [-][-]")
		log.Fatalln(err)
	}

}
