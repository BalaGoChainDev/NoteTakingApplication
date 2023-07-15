package utils

import (
	// Standard libs
	"fmt"
	"net/http"
	"strings"

	// ThirdParty libs
	"github.com/go-playground/validator/v10"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
)

// ValidateInput validates the input request data using the validator library.
// It returns an ErrorResponder in case of validation errors, containing the error message and status code.
func ValidateInput(reqData interface{}) response.ErrorResponder {
	v := validator.New()
	if err := v.Struct(reqData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0, len(validationErrors))
		for _, e := range validationErrors {
			fieldName := e.Field()
			switch e.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is required", fieldName))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' must be a valid email address", fieldName))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' must be at least %s characters long", fieldName, e.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' cannot exceed %s characters", fieldName, e.Param()))
			default:
				errorMessages = append(errorMessages, e.Error())
			}
		}
		errorMessage := strings.Join(errorMessages, ", ")
		return response.NewErrorResponse(http.StatusBadRequest, errorMessage)
	}

	return nil
}
