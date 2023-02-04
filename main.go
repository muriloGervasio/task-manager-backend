package main

import (
	"app/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/:project/percentage", handlers.TBpercentageHandler)
	app.Get("/:project/tasks", handlers.GetTasks)

	log.Fatal(app.Listen(":3000"))

}
