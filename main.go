package main

import (
	"app/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("i am alive...")
	})
	app.Get("/:project/percentage", handlers.TBpercentageHandler)
	app.Get("/:project/tasks", handlers.GetTasks)

	log.Fatal(app.Listen(":8080"))
}
