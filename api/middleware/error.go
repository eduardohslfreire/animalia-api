package middleware

import (
	"errors"
	"strings"

	http_errors "github.com/eduardohslfreire/animalia-api/api/errors"
	"github.com/eduardohslfreire/animalia-api/api/validation"
	business_errors "github.com/eduardohslfreire/animalia-api/entity/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ErrorMiddleware ...
func (m *Middleware) ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		errResponse := buildErrorResponse(c.Errors[0].Err)
		c.JSON(errResponse.StatusCode(), errResponse)
	}
}

func buildErrorResponse(err error) http_errors.ErrorResponse {
	switch err.(type) {
	case validator.ValidationErrors:
		return http_errors.BadRequest(extractMessageValidationErrors(err))
	case business_errors.BusinessError:
		return http_errors.UnprocessableEntity(err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http_errors.NotFound("")
	}

	return http_errors.InternalServerError("")
}

func extractMessageValidationErrors(err error) string {
	errors := err.(validator.ValidationErrors)
	messages := errors.Translate(validation.Translator)
	message := make([]string, 0, len(messages))

	for _, msg := range messages {
		message = append(message, msg)
	}

	return strings.Join(message, ", ")
}
