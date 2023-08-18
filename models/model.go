package models

import (
	"time"

	"gorm.io/gorm"
)

type ModelID struct {
	ID uint `gorm:"primarykey" json:"id"`
}

type ModelTimeStamps struct {
	CreatedAt time.Time `gorm:"primarykey" json:"created_at"`
	UpdatedAt time.Time `gorm:"primarykey" json:"updated_at"`
}

type ModelTimeStampsWithDeletedAt struct {
	CreatedAt time.Time      `gorm:"primarykey" json:"created_at"`
	UpdatedAt time.Time      `gorm:"primarykey" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
