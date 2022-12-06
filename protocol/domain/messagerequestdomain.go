package domain

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type MessageRequestDomain struct {
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
	Message    string `json:"message"`
}

func (m *MessageRequestDomain) Validate() []string {
	var errors []string
	message := *m
	err := validation.ValidateStruct(
		&message,
		validation.Field(&message.SenderId, validation.Required, is.UUIDv4),
		validation.Field(&message.ReceiverId, validation.Required, is.UUIDv4),
		validation.Field(&message.Message, validation.Required),
	)
	if err != nil {
		errors = append(errors, strings.Split(err.Error(), "; ")...)
	}

	return errors
}
