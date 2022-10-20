package dto

import "time"

type SocialmediaEditDto struct {
	SocialMediaId  int    	  `json:"Social_media_id"`
	Name           string 	  `json:"name"`
	SocialMediaUrl string 	  `json:"social_media_url"`
	UserId         int	  	  `json:"user_id"`
	UpdatedAt 	   *time.Time `json:"updated_at,omitempty"`
}