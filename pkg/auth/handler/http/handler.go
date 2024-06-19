package http

import (
	"api-project/internal/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

/*
handler is a struct that holds the methods to handle authentication-related HTTP requests. It
encapsulates the business logic defined in the auth.Usecase interface.
*/
type handler struct {
	usecase interfaces.AuthUsecase
}

// New initializes and configures the authentication routes for the application.
func New(app fiber.Router, usecase interfaces.AuthUsecase) {
	handler := &handler{
		usecase: usecase,
	}

	app.Post("login/", handler.LoginHandler())
	app.Post("logout/", handler.LogoutHandler())
	app.Post("create-access-token/", handler.GenerateAccessFromRefreshHandler())

	app.Post("register/", handler.RegisterUser())

}
