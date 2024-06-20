package main

import (
	"log"

	"api-project/internal/config"
	"api-project/internal/domain/interfaces"

	jwtWare "github.com/gofiber/jwt/v2"

	_authHandler "api-project/pkg/auth/handler/http"
	"api-project/pkg/auth/handler/middleware.go"
	_authRepo "api-project/pkg/auth/repository/gorm"
	_authUsecase "api-project/pkg/auth/usecase"

	_transactionHandler "api-project/pkg/transaction/handler/http"
	_transactionRepo "api-project/pkg/transaction/repository/gorm"
	_transactionUsecase "api-project/pkg/transaction/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type server struct {
	// Pointer to the GORM database client used by the server
	DB *gorm.DB
	// Pointer to the Fiber application instance used by the server
	App *fiber.App
}
type usecases struct {
	transactionUsecase interfaces.TransactionUsecase
	authUsecase        interfaces.AuthUsecase
}

func main() {

	// Initiate viper configuration for reading config file
	config.ConfigureViper()

	// Initialize global server configuration parameters
	config.InitServerConfig()

	// Call global server config parameter variable
	server := config.Server

	app := server.App

	// Initialize a connection to the database
	db := config.InitDB(viper.GetBool("verbose"), viper.GetBool("logger"))
	if db != nil {
		log.Println("[+] Success: Connection to database successful [+]")
	}

	// Logger middleware with configuration specified in logger.Config{} to log requests to the server
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path} ${time}\n",
	}))

	// Middleware for caching, compression, and logging in the order specified
	app.Use(
		logger.New(config.LoggerConfig()),
		compress.New(config.CompressResponseConfig()),
	)

	// CORS middleware to allow cross-origin requests
	server.App.Use(cors.New(config.CorsConfig()))

	// Configure authentication and authorization middleware
	authMW := middleware.Protected(viper.GetString("jwt.access_secret"))
	jwtMW := middleware.ValidateJWT(server.DB, server.RedisClient, jwtWare.Config{})

	/* authRoute is a Fiber.Router group created under the "/auth/" path, to which authentication
	related routes will be added */
	authRoute := app.Group("/auth/")

	/* protectedRoutes is a Fiber.Router group created under the "/api/" path, to which other
	private routes will be added */
	protectedRoutes := server.App.Group("/api/", authMW, jwtMW)

	// Initialize auth repository and usecase
	authRepo := _authRepo.New(server.DB)

	// Initialize transaction repository and usecase
	transactionRepo := _transactionRepo.New(server.DB)

	// Initialize usecases struct with required usecase instances
	usecases := &usecases{
		transactionUsecase: _transactionUsecase.New(transactionRepo, authRepo),
		authUsecase:        _authUsecase.New(authRepo, transactionRepo, server.RedisClient),
	}

	initRoutes(authRoute, protectedRoutes, usecases)

	// Start the Fiber web server on the specified port
	log.Fatal(app.Listen(":" + viper.GetString(`server.port`)))
}

/*
initRoutes initializes routes for the all endpoints on the provided Fiber router. It sets up the
necessary handlers and routes for handling operations.
*/
func initRoutes(authRoute fiber.Router, protectedRoutes fiber.Router, usecases *usecases) {
	_authHandler.New(authRoute, usecases.authUsecase)
	_transactionHandler.New(protectedRoutes, usecases.transactionUsecase)
}
