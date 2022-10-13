package models

type Drink struct {
	GormModel
	Water 		int
	WaterStatus string
	Wish		int
	WishStatus	string
}