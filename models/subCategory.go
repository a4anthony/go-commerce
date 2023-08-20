package models

import "database/sql"

type SubCategory struct {
	ModelID
	Name        string       `gorm:"size:255;not null;" json:"name"`
	Slug        string       `gorm:"size:255;not null;" json:"slug"`
	Description string       `gorm:"size:255;not null;" json:"description"`
	IsActive    sql.NullBool `gorm:"default:false" json:"is_active"`
	CategoryID  uint         `gorm:"not null" json:"category_id"`
	Category    Category     `gorm:"foreignKey:CategoryID" json:"category"`
	ModelTimeStamps
}
