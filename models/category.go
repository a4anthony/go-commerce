package models

type Category struct {
	ModelID
	Name        string        `gorm:"size:255;not null;" json:"name"`
	Slug        string        `gorm:"size:255;not null;unique" json:"slug"`
	Description string        `gorm:"size:255;not null;" json:"description"`
	IsActive    bool          `gorm:"default:false" json:"is_active"`
	SubCategory []SubCategory `gorm:"foreignKey:CategoryID" json:"sub_categories"`
	ModelTimeStamps
}
