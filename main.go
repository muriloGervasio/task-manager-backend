package main

import (
	"app/config"
	"app/handlers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("i am alive...")
	})
	app.Get("/:project/percentage", handlers.TBpercentageHandler)
	app.Get("/:project/tasks", handlers.GetTasks)
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "not found")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config.Port)))
}
