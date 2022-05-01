package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model

	Content string `gorm:"type:VARCHAR(255);not null;validate:required,gte=0,lte=130"`
	RoomID int 
}
