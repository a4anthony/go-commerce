package models

type SubCategory struct {
	ModelID
	Name        string   `gorm:"size:255;not null;" json:"name"`
	Slug        string   `gorm:"size:255;not null;" json:"slug"`
	Description string   `gorm:"size:255;not null;" json:"description"`
	IsActive    bool     `gorm:"default:false" json:"is_active"`
	CategoryID  uint     `gorm:"not null" json:"-"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ModelTimeStamps
}
