package errors

import (
	"encoding/json"
	"fmt"
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/go-playground/validator"
	"strings"
)

func ValidationError(bindError error) *response.BaseResponse {
	switch errorType := bindError.(type) {
	case *json.SyntaxError:
		message := "invalid json"
		return response.NewBadRequestError(message)
	case *json.UnmarshalTypeError:
		message := fmt.Sprintf("%s is invalid", errorType.Field)
		return response.NewBadRequestError(message)
	case validator.ValidationErrors:
		for _, validationError := range errorType {
			validationMessage := fmt.Sprintf("%s is a required field", strings.ToLower(validationError.Field()))
			return response.NewBadRequestError(validationMessage)
		}
	default:
		message := "Cannot process fields. Please check and try again"
		return response.NewBadRequestError(message)
	}

	return nil
}
