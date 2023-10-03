package helper

import "github.com/gofiber/fiber/v2"

func JSONIfError(err error, c *fiber.Ctx, m string) error {
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"status":  false,
		"error":   err.Error(),
		"message": m,
	})
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}