package dto

import "time"
type CommentUserDto struct {
	UserId   int 	    `json:"user_id"`
	Username string		`json:"username"`
	Email    string		`json:"email"`
}

type CommentPhotoDto struct {
	PhotoId  int         `json:"photo_id"`
	Title    string      `json:"title"`
	Caption  string      `json:"caption"`
	PhotoUrl string      `json:"photo_url" `
	UserId   int         `json:"user_url"`
}

type CommentAllDto struct {
	CommentId int    	 `json:"comment_id"`
	Message   string 	 `json:"message"`
	UserId    int    	 `json:"user_id"`
	PhotoId   int    	 `json:"photo_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	User 	  CommentUserDto
	Photo 	  CommentPhotoDto
}