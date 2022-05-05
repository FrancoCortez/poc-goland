package entity

import (
	"gorm.io/gorm"
	"time"
)

type Attribute struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:50;"`
}
