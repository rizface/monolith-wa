package entity

import "time"

type Message struct {
	Id         string    `json:"id"`
	SenderId   string    `json:"senderId"`
	ReceiverId string    `json:"receiverId"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
