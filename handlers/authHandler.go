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

type LoginInput struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=32" json:"password"`
}

type UserResponse struct {
	Message     string      `json:"message"`
	User        models.User `json:"user,omitempty"`
	UserJSON    []byte      `json:"user_json,omitempty"`
	AccessToken string      `json:"access_token,omitempty"`
}

var validate = validator.New()

// @Summary Regiuster new user.
func Register(c *fiber.Ctx) error {
	user := new(RegisterInput)
	if err := c.BodyParser(user); err != nil {
		if err.Error() == "Unprocessable Entity" {
			user.Email = ""
			user.Password = ""
			user.Phone = ""
			user.FirstName = ""
			user.LastName = ""
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

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

	// check if user was deleted
	deletedUser := &models.User{}
	database.DB.Unscoped().Where("email = ? AND deleted_at IS NOT NULL", user.Email).First(&deletedUser)

	if deletedUser != nil {
		database.DB.Unscoped().Delete(&deletedUser)
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

func Login(c *fiber.Ctx) error {
	user := new(LoginInput)

	if err := c.BodyParser(user); err != nil {
		if err.Error() == "Unprocessable Entity" {
			user.Email = ""
			user.Password = ""
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	validationErr := validate.Struct(*user)
	if validationErr != nil {
		fmt.Println(user)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.SetError(validationErr))
	}

	token, err := models.LoginCheck(user.Email, user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.ErrorMsg{
				Message: "Validation error",
				Errors: map[string]string{
					"error": "Invalid credentials.",
				},
			},
		)
	}

	authenticatedUser := models.User{}
	database.DB.Where("email = ?", user.Email).First(&authenticatedUser)

	return c.Status(fiber.StatusCreated).JSON(
		UserResponse{
			User:        authenticatedUser,
			AccessToken: token,
			Message:     "User logged in successfully.",
		},
	)
}

func Me(c *fiber.Ctx) error {
	uID, err := utils.ExtractTokenID(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.ErrorMsg{
				Message: "Validation error",
				Errors: map[string]string{
					"error": "Invalid credentials.",
				},
			},
		)
	}

	user := models.User{}
	database.DB.Where("id = ?", uID).First(&user)

	// check if user exists
	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.ErrorMsg{
				Message: "Validation error",
				Errors: map[string]string{
					"error": "Invalid credentials.",
				},
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		UserResponse{
			User:    user,
			Message: "User retrieved successfully.",
		},
	)

}

func DeleteUser(c *fiber.Ctx) error {
	uID, err := utils.ExtractTokenID(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.ErrorMsg{
				Message: "Validation error",
				Errors: map[string]string{
					"error": "Invalid credentials.",
				},
			},
		)
	}

	user := models.User{}
	database.DB.Where("id = ?", uID).First(&user)

	// check if user exists
	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.ErrorMsg{
				Message: "Validation error",
				Errors: map[string]string{
					"error": "Invalid credentials.",
				},
			},
		)
	}

	database.DB.Delete(&user)

	return c.Status(fiber.StatusOK).JSON(
		UserResponse{
			User:    user,
			Message: "User deleted successfully.",
		},
	)

}
