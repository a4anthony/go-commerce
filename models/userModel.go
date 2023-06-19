package models

import (
	"database/sql"
	"encoding/json"

	"github.com/a4anthony/go-commerce/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ModelID
	FirstName       string   `gorm:"size:255;not null;" json:"first_name"`
	LastName        string   `gorm:"size:255;not null;" json:"last_name"`
	Phone           string   `gorm:"size:255;not null;" json:"phone"`
	Email           string   `gorm:"size:255;not null;unique" json:"email"`
	Password        string   `gorm:"size:255;not null;" json:"-"`
	EmailVerifiedAt NullTime `gorm:"index, null" json:"email_verified_at"`
	ModelTimeStamps
}

type NullTime struct {
	sql.NullTime
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	}
	return []byte("null"), nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil

}

func (u *User) SaveUser() (*User, error) {
	err := database.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
