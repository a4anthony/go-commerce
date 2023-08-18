package models

import "database/sql"

type Category struct {
	ModelID
	Name        string       `gorm:"size:255;not null;" json:"name"`
	Slug        string       `gorm:"size:255;not null;unique" json:"slug"`
	Description string       `gorm:"size:255;not null;" json:"description"`
	IsActive    sql.NullBool `gorm:"default:false" json:"is_active"`
	ModelTimeStamps
}
