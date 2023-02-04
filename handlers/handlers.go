package handlers

import (
	"app/connection"
	useCase "app/use-case"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func TBpercentageHandler(c *fiber.Ctx) error {
	percentage := useCase.GetActiveTaskPercentage(*connection.Client, c.Params("project"))

	return c.JSON(fiber.Map{
		"percentage": fmt.Sprintf("%f", percentage),
	})
}

func GetTasks(c *fiber.Ctx) error {
	tasks := useCase.GetTdTasks(*connection.Client, c.Params("project"))

	return c.JSON(tasks)
}
