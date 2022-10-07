package models


type Item struct {
	ItemId   	int			`gorm:"primaryKey;autoIncrement;"` 
	ItemCode 	string		 
	Desription 	string		 
	Quantity	int			 
	OrderId		int			 `gorm:"references:OrderId"`
}