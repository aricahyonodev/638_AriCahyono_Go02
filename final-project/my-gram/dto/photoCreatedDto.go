package dto

import "time"

type PhotoCreateDto struct {
	PhotoId  int         `json:"photo_id"`
	Title    string      `json:"title"`
	Caption  string      `json:"caption"`
	PhotoUrl string      `json:"photo_url" `
	UserId   int         `json:"user_url"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}