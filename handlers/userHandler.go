package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mnurichsan/golang-rest-api/database"
	"github.com/mnurichsan/golang-rest-api/models"
	"gorm.io/gorm"
)

type User models.User

func GetUserIndex(c *fiber.Ctx) error {
	db := database.DB
	user := &User{}
	err := db.First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Get User",
		"data":    user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB

	type CreateUserRequest struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
	}

	requestUser := new(CreateUserRequest)

	if err := c.BodyParser(requestUser); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"data":    err,
		})
	}

	newUser := User{
		Name:      requestUser.Name,
		Email:     requestUser.Email,
		Password:  requestUser.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if resultError := db.Create(&newUser).Error; resultError != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Invalid Create Data",
			"data":    resultError,
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Ok",
		"data":    newUser,
	})

}
