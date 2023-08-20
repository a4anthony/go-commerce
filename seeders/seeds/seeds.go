package seeds

import (
	"fmt"

	"github.com/a4anthony/go-commerce/database"
	"github.com/a4anthony/go-commerce/seeders"
	"gorm.io/gorm"
)

func All() []seeders.Seed {
	output := []seeders.Seed{}
	output = append(output, userSeed()...)
	output = append(output, categorySeed()...)
	return output
}

func userSeed() []seeders.Seed {
	fmt.Println("Creating users...")
	return []seeders.Seed{
		{
			Name: "CreateA4anthony",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "Anthony", "Akro", "0700000000")
				return err
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "John", "Doe", "0700000001")
				return err
			},
		},
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "Jane", "Doe", "0700000002")
				return err
			},
		},
		{
			Name: "CreateSusan",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "Susan", "Doe", "0700000003")
				return err
			},
		},
	}
}

func categorySeed() []seeders.Seed {
	fmt.Println("Creating categories...")

	// Create categories with sub_categories array
	categories := map[string]interface{}{
		"Electronics": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Home": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Kitchen": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Phones": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Computers": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Groceries": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Beauty": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Health": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Automobile": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
			"Sub category 4",
			"Sub category 5",
		},
		"Books": []string{
			"Sub category 1",
			"Sub category 2",
			"Sub category 3",
		},
		"Baby": []string{
			"Sub category 1",
			"Sub category 2",
		},
	}

	db := database.DB

	for category, subCategories := range categories {
		fmt.Println(category, subCategories)
		runFunc := func(db *gorm.DB) error {
			category, err := CreateCategory(db, category, category+" category description")
			if category.ID > 0 {
				for _, subCategory := range subCategories.([]string) {
					CreateSubCategory(db, subCategory, subCategory+" sub category description", category.ID)
				}
			}
			return err
		}
		runFunc(db)
	}

	output := []seeders.Seed{}
	return output
}
