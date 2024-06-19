package config

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// server represents an instance of our application server
type server struct {
	// Pointer to the GORM database client used by the server
	DB *gorm.DB

	// Pointer to the Redis client used by the server
	RedisClient *redis.Client

	// Pointer to the Fiber application instance used by the server
	App *fiber.App
}

/*
Server is a global variable initialized with database, redis client and fiber app required for the
application
*/
var Server = &server{}

/*
InitServerConfig initializes the configuration parameters for the server with database, redis client
and fiber app globally.
*/
func InitServerConfig() {
	// Initialize a connection to the database
	db := InitDB(viper.GetBool("verbose"), viper.GetBool("logger"))
	if db == nil {
		log.Fatal("[-] Error: Error connecting to the database [-]")
	}

	// Initialize a connection to the redis server
	redisClient, err := ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize server sruct with the database instance
	Server = &server{
		DB:          db,
		RedisClient: redisClient,
		App:         fiber.New(NewFiberConfig()),
	}
}
