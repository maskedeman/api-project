package utils

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

/*
InitTranslator initializes the validator and translator instances for the application. It returns the validator and translator instances.
The validator instance is used to validate request payloads, while the translator instance is used to translate validation errors into
the desired language.

The default language is English.
*/
func InitTranslator() (*validator.Validate, ut.Translator) {

	// Create a new validator instance
	validate := validator.New()

	// Initialize English language support for translation
	en := en.New()

	// Create a new universal translator with English as the default language
	uni := ut.New(en, en)

	// Get the translator instance for the English language
	trans, _ := uni.GetTranslator("en")

	// Register default translations for English in the validator
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	// Return the validator and translator instances
	return validate, trans
}

func TranslateError(validationErrs validator.ValidationErrors, trans ut.Translator) map[string]string {

	// Create an empty map to store field names as key and their translated error messages as value
	err_map := make(map[string]string)

	// Check if validation errors are provided
	if len(validationErrs) != 0 {
		for _, e := range validationErrs {
			// Store the lowercase field name and its translated error message in the map
			err_map[strings.ToLower(e.Field())] = e.Translate(trans)
		}
	}

	return err_map
}
