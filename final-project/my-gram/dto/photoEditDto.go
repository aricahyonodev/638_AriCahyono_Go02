package dto

import "time"

type PhotoEditDto struct {
	PhotoId  int         `json:"photo_id"`
	Title    string      `json:"title"`
	Caption  string      `json:"caption"`
	PhotoUrl string      `json:"photo_url" `
	UserId   int         `json:"user_url"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}