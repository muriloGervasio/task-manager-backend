package main

import (
	"app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/tb-percentage", handlers.TBpercentageHandler)
	app.Get("/gcl-percentage", handlers.GCLpercentageHandler)

	app.Listen(":3000")

}
