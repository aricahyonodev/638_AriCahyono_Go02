package response

import "my-gram/models"

type UserComment struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type PhotoComment struct {
	PhotoId  int    `json:"photo_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type ResponseComment struct {
	CommentId int    `gorm:"primaryKey;autoIncrement" json:"comment_id"`
	Message   string `json:"message" valid:"required~Message wajib terisi"`
	UserId    int    `gorm:"references:UserId" json:"user_id"`
	PhotoId   int    `gorm:"references:PhotoId" json:"photo_id"`
	User      UserComment
	Photo     PhotoComment
	models.GormModel
}