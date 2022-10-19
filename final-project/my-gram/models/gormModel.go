package models

import "time"

type GormModel struct {
	CreateAt *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdateAt *time.Time `gorm:"autoCreateTime" json:"updated_at,omitempty"`
}