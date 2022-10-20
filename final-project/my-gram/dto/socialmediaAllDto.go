package dto

import "time"
type SocialmediaUserDto struct {
	UserId   		   int 	    `json:"user_id"`
	Username 		   string	`json:"username"`
	ProfileImageUrl    string	`json:"profile_image_url"`
}


type SocialmediaAllDto struct {
	SocialMediaId  int    	  `json:"Social_media_id"`
	Name           string 	  `json:"name"`
	SocialMediaUrl string 	  `json:"social_media_url"`
	UserId         int	  	  `json:"user_id"`
	CreatedAt 	   *time.Time `json:"created_at,omitempty"`
	UpdatedAt 	   *time.Time `json:"updated_at,omitempty"`

	User 	  SocialmediaUserDto
}