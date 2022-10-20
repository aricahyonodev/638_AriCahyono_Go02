package dto

type UserDto struct {
	UserId   int 	    	`json:"user_id"`
	Username string			`json:"username"`
	Email    string			`json:"email"`
	Age      int			`json:"age"`	
	ProfileImageUrl string	`json:"profile_image_url"`

}