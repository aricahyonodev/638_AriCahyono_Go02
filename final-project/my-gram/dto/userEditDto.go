package dto

import "time"

type UserEditDto struct {
	UserId   int         `json:"user_id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Age      int         `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}