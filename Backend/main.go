package main

import (
	"backend/Database"
	"backend/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New(fiber.Config{
		//	Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Hesat Backend",
	})

	if fiber.IsChild() == false {
		Database.Init()
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://127.0.0.1:8000, http://127.0.0.1:8000, http://localhost:8000, http://localhost:5173, http://localhost:4173",
		AllowHeaders:     "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))

	app.Static("/", "./public")
	Routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
