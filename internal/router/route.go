package router
import (
	"laundry-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	User *handler.UserHandler
}

func SetupRoutes(app *fiber.App, h *Handlers) {

	users := app.Group("/users")
	users.Post("/", h.User.CreateUser)
	app.Get("/users/role/:role", h.User.GetUserByRole)
	app.Get("/users/phone/:phone_number", h.User.GetUserByPhone)
	app.Get("/users", h.User.GetAllUsers)
	users.Put("/:id", h.User.UpdateUser)
}