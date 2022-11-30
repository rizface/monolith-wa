package constant

import "net/http"

type ErrorBuilder struct {
	Message     string
	Code        int
	RealMessage string
}

var (
	USERNAME_IS_EXISTS = &ErrorBuilder{
		Message: "USERNAME_IS_EXISTS",
		Code:    http.StatusUnprocessableEntity,
	}
	USER_NOT_FOUND = &ErrorBuilder{
		Message: "USER_NOT_FOUND",
		Code:    http.StatusNotFound,
	}
	PASSWORD_WRONG = &ErrorBuilder{
		Message: "PASSWORD_WRONG",
		Code:    http.StatusUnauthorized,
	}
)

func InternalServerError(msg string) *ErrorBuilder {
	return &ErrorBuilder{
		Message:     "INTERNAL_SERVER_ERROR",
		Code:        http.StatusInternalServerError,
		RealMessage: msg,
	}
}
