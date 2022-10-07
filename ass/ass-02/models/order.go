package models

import (
	"time"
)

type Order struct {
	OrderId      int        `gorm:"primaryKey;autoIncrement;"`
	CustomerName string		
	Items 		 []Item	 	
	OrderedAt      time.Time 
}