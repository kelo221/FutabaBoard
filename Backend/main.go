package main

import (
	"backend/Database"
	"backend/Routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	app := fiber.New()

	if fiber.IsChild() == false {
		Database.Init()
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173, http://localhost:4173",
		AllowHeaders:     "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))

	app.Static("/", "./public")
	Routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
