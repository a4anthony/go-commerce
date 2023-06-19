package handlers

import (
	"fmt"

	"github.com/a4anthony/go-commerce/database"
	"github.com/a4anthony/go-commerce/models"
	"github.com/a4anthony/go-commerce/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterInput struct {
	FirstName string `validate:"required,min=2,max=32" json:"first_name"`
	LastName  string `validate:"required,min=2,max=32" json:"last_name"`
	Phone     string `validate:"required,numeric" json:"phone"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,min=6,max=32" json:"password"`
}

type UserResponse struct {
	Message     string      `json:"message"`
	User        models.User `json:"user,omitempty"`
	UserJSON    []byte      `json:"user_json,omitempty"`
	AccessToken string      `json:"access_token,omitempty"`
}

var validate = validator.New()

func Register(c *fiber.Ctx) error {
	user := new(RegisterInput)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validate.Struct(*user)
	if validationErr != nil {
		fmt.Println(user)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.SetError(validationErr))
	}

	// check if email already exists
	emailExists := database.DB.Where("email = ?", user.Email).First(&models.User{}).Error
	if emailExists == nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			utils.ErrorMsg{
				Message: "Validation error",
				Errors: map[string]string{
					"email": "The email has already been taken.",
				},
			},
		)
	}

	u := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
	}

	_, err := u.SaveUser()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createdUser := models.User{}
	database.DB.Where("email = ?", user.Email).First(&createdUser)

	return c.Status(fiber.StatusCreated).JSON(UserResponse{
		User:        createdUser,
		AccessToken: token,
		Message:     "User created successfully.",
	})

}
