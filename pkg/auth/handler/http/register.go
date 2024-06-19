package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-project/internal/domain/presenters"

	"api-project/internal/domain/models"
	"api-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
Register is a Fiber handler function for user registration. It parses the request JSON body,
validates the user information, and invokes the use case to register the user. If any validation
or registration error occurs, it returns an appropriate error response. Otherwise, it returns a
success response
*/
func (handler *handler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenters.UserCreateRequest
		var user models.User

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		data, err := json.Marshal(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}
		fmt.Printf("user: %v\n", requestBody.InitialDeposit)
		errMap = handler.usecase.RegisterUser(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		return c.JSON(presenters.SuccessResponse())
	}
}
