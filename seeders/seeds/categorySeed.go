package seeds

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/a4anthony/go-commerce/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name, description string) (models.Category, error) {

	// this is type NullBool struct {
	// 	sql.NullBool
	// }

	// status := models.NullBool{}
	// set status to true or false randomly

	//
	//

	// status := models.NullBool{}
	// if rand.Intn(2) == 1 {
	// 	status = models.NullBool{Bool: true, Valid: true}
	// } else {
	// 	status = sql.NullBool{Bool: false, Valid: true}
	// }

	status := rand.Intn(2) == 1

	// check if category with is_active = true exceeds 5
	var count int64
	db.Model(&models.Category{}).Where("is_active = ?", false).Count(&count)
	if count >= 2 {
		status = true
	}

	category := models.Category{
		Name:        name,
		Slug:        strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		Description: description,
		IsActive:    status,
	}

	categoryMap := map[string]interface{}{
		"name":        name,
		"slug":        strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		"description": description,
		"is_active":   status,
	}

	// create from map
	err := db.Model(&models.Category{}).Create(categoryMap).Error

	// retrieve from category by slug from db
	db.Model(&models.Category{}).Where("slug = ?", category.Slug).First(&category)
	return category, err

	// fmt.Println(categoryMap)
	// err := db.Create(&category).Error
	// return category, err
}

func CreateSubCategory(db *gorm.DB, name, description string, categoryID uint) (models.SubCategory, error) {
	fmt.Println("subCategory")

	status := rand.Intn(2) == 1

	// check if category with is_active = true exceeds 5
	var count int64
	db.Model(&models.SubCategory{}).Where("is_active = ?", false).Count(&count)
	if count >= 2 {
		status = true
	}

	subCategory := models.SubCategory{
		Name:        name,
		Slug:        strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		Description: description,
		IsActive:    status,
		CategoryID:  categoryID,
	}

	fmt.Println("subCategory")
	err := db.Create(&subCategory).Error
	return subCategory, err
}
