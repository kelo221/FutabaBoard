package Routes

import (
	"backend/Controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")
	admin.Delete("remove", Controllers.DeletePost)
	admin.Post("ban", Controllers.Ban)

	api.Post("register", Controllers.Register)
	api.Post("login", Controllers.Login)
	api.Post("thread", Controllers.Thread)
	api.Post("post", Controllers.Post)

	api.Get("thread/*", Controllers.FetchThread)
	api.Get("post/*", Controllers.FetchPost)
	api.Get("test", Controllers.TestFunction)

}
