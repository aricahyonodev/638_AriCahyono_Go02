package dto

import "time"

type CommentEditDto struct {
	CommentId int    	 `json:"comment_id"`
	Message   string 	 `json:"message"`
	UserId    int    	 `json:"user_id"`
	PhotoId   int    	 `json:"photo_id"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}