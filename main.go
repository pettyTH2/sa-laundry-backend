package main

import (
	"laundry-backend/internal/database"
	"laundry-backend/internal/handler"
	"laundry-backend/internal/repository/postgres"
	"laundry-backend/internal/usecase"
	"laundry-backend/internal/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.NewDatabaseConnection()

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}

	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(*userUsecase)

	app := fiber.New()

	handlers := &routes.Handlers{
		User: userHandler,
	}
	
	routes.SetupRoutes(app, handlers)
	log.Fatal(app.Listen(":3000"))
}