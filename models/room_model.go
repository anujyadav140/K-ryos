package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(255);not null;uniqueIndex" validate:"required,gte=3,lte=30"`
	Desc string `gorm:"type:VARCHAR(255);not null" validate:"required,gte=10,lte=100"`
	Messages []Message
}