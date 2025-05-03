package main

import (
	"log"
	"yaltopia_task/data"
	"yaltopia_task/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	if err := data.Parse(); err != nil {
		log.Fatalf("parsing error: %s", err.Error())
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	cricket := app.Group("/cricket")
	routes.RegisterCricketRoutes(cricket)

	volleyball := app.Group("/volleyball")
	routes.RegisterVolleyballRoutes(volleyball)

	app.Listen(":7777")
}
