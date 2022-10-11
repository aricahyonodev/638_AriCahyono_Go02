package models

import "time"

type GormModel struct {
	ID       uint       `gorm:"primaryKey" jsno:"id"`
	CreateAt *time.Time `json:"created_at,omitempty"`
	UpdateAt *time.Time `json:"updated_at,omitempty"`
}