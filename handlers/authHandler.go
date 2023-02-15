package handlers

import "github.com/gofiber/fiber/v2"

type RequestUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	request := new(RequestUser)

	if err := c.BodyParser(request); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"data":    err,
		})
	}

	user := &User{
		Name:     request.Name,
		Password: request.Password,
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Test Request",
		"data":    user,
	})

}
