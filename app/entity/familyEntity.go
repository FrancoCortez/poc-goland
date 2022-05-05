package entity

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Families []*Family

type Family struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `gorm:"size:50;unique"`
	Status         string         `gorm:"size:50;"`
	FamilyGroup    FamilyGroup
	FamilyExamples []FamilyExample `gorm:"foreignKey:Family"`
}

type FamilyDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func (b Family) ToDto() *FamilyDto {
	return &FamilyDto{
		ID:     b.ID,
		Name:   b.Name,
		Status: b.Status,
	}
}

type FamilyDtos []*FamilyDto

func (bs Families) ToDto() FamilyDtos {
	dtos := make([]*FamilyDto, len(bs))
	for i, b := range bs {
		fmt.Println(b.FamilyGroup)
		dtos[i] = b.ToDto()
	}
	return dtos
}
