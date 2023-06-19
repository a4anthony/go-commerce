package seeds

import (
	"strings"

	"github.com/a4anthony/go-commerce/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, fname string, lname string, phone string) error {
	u := models.User{}

	u.FirstName = fname
	u.LastName = lname
	u.Phone = phone
	u.Password = "password"
	u.Email = strings.ToLower(fname) + "@email.com"
	// u.EmailVerifiedAt = sql.NullTime{}
	if db.Where("email = ?", u.Email).First(&u).Error == nil {
		return nil
	}

	return db.Create(&u).Error
}
