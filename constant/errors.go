package constant

import "net/http"

type ErrorBuilder struct {
	Message     string `json:"message"`
	Code        int    `json:"code"`
	RealMessage string `json:"realMessage"`
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
	INVALID_TOKEN = &ErrorBuilder{
		Message: "INVALID_JWT",
		Code:    http.StatusUnauthorized,
	}
	TOKEN_EXPIRED = &ErrorBuilder{
		Message: "JWT_TOKEN_EXPIRED",
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
