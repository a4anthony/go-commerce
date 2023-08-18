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

	categories := []string{
		"Electronics",
		"Home",
		"Kitchen",
		"Phones",
		"Computers",
		"Groceries",
		"Beauty",
		"Health",
		"Automobile",
		"Books",
		"Baby",
	}
	output := []seeders.Seed{}
	db := database.DB
	for _, category := range categories {
		runFunc := func(db *gorm.DB) error {
			err := CreateCategory(db, category, category+" category description")
			return err
		}
		runFunc(db)

	}
	return output
}
