package dto

import "time"

type CommentCreatedDto struct {
	CommentId int    	 `json:"comment_id"`
	Message   string 	 `json:"message"`
	UserId    int    	 `json:"user_id"`
	PhotoId   int    	 `json:"photo_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}