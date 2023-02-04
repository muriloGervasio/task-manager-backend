package handlers

import (
	"app/connection"
	useCase "app/use-case"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func TBpercentageHandler(c *fiber.Ctx) error {
	percentage := useCase.GetActiveTaskPercentage(*connection.Client, "TB")

	return c.JSON(fiber.Map{
		"percentage": fmt.Sprintf("%f", percentage),
	})
}

func GCLpercentageHandler(c *fiber.Ctx) error {
	percentage := useCase.GetActiveTaskPercentage(*connection.Client, "GCL")

	return c.JSON(fiber.Map{
		"percentage": fmt.Sprintf("%f", percentage),
	})
}
