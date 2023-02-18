package model

import "time"

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}
