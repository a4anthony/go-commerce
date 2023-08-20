package seeds

import (
	"database/sql"
	"math/rand"
	"strings"

	"github.com/a4anthony/go-commerce/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name, description string) (models.Category, error) {

	status := sql.NullBool{}
	if rand.Intn(2) == 1 {
		status = sql.NullBool{Bool: true, Valid: true}
	} else {
		status = sql.NullBool{Bool: false, Valid: true}
	}

	// check if category with is_active = true exceeds 5
	var count int64
	db.Model(&models.Category{}).Where("is_active = ?", false).Count(&count)
	if count >= 2 {
		status = sql.NullBool{Bool: true, Valid: true}
	}

	category := models.Category{
		Name:        name,
		Slug:        strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		Description: description,
		IsActive:    status,
	}

	err := db.Create(&category).Error
	return category, err
}

func CreateSubCategory(db *gorm.DB, name, description string, categoryID uint) (models.SubCategory, error) {

	status := sql.NullBool{}
	if rand.Intn(2) == 1 {
		status = sql.NullBool{Bool: true, Valid: true}
	} else {
		status = sql.NullBool{Bool: false, Valid: true}
	}

	// check if category with is_active = true exceeds 5
	var count int64
	db.Model(&models.SubCategory{}).Where("is_active = ?", false).Count(&count)
	if count >= 2 {
		status = sql.NullBool{Bool: true, Valid: true}
	}

	subCategory := models.SubCategory{
		Name:        name,
		Slug:        strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		Description: description,
		IsActive:    status,
		CategoryID:  categoryID,
	}

	err := db.Create(&subCategory).Error
	return subCategory, err
}
