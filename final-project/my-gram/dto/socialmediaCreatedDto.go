package dto

import "time"

type SocialmediaCreatedDto struct {
	SocialMediaId  int    	  `json:"Social_media_id"`
	Name           string 	  `json:"name"`
	SocialMediaUrl string 	  `json:"social_media_url"`
	UserId         int	  	  `json:"user_id"`
	CreatedAt 	   *time.Time `json:"created_at,omitempty"`

}