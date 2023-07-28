package main

import (
	"backend/Database"
	"backend/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	Database.Init()
	app := fiber.New()

	app.Static("/", "./public")

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://127.0.0.1:8000, http://127.0.0.1:8000, http://localhost:8000, http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	Routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
