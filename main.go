package main

import (
	"laundry-backend/internal/config"
	"laundry-backend/internal/database"
	"laundry-backend/internal/handler"
	"laundry-backend/internal/repository/postgres"
	"laundry-backend/internal/router"
	"laundry-backend/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize database
	db := database.NewDatabaseConnection()

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}

	userRepo := postgres.NewUserRepository(db)
	couponRepo := postgres.NewCouponRepository(db)
	userCouponRepo := postgres.NewUserCouponRepository(db)
	clothRepo := postgres.NewClothRepository(db)
	clothListRepo := postgres.NewClothListRepository(db)
	orderRepo := postgres.NewOrderRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	couponUsecase := usecase.NewCouponUsecase(couponRepo)
	userCouponUsecase := usecase.NewUserCouponUsecase(userCouponRepo)
	clothUsecase := usecase.NewClothUsecase(clothRepo)
	clothListUsecase := usecase.NewClothListUsecase(clothListRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)

	userHandler := handler.NewUserHandler(*userUsecase)
	couponHandler := handler.NewCouponHandler(*couponUsecase)
	userCouponHandler := handler.NewUserCouponHandler(*userCouponUsecase)
	clothHandler := handler.NewClothHandler(*clothUsecase)
	clothListHandler := handler.NewClothListHandler(*clothListUsecase)
	orderHandler := handler.NewOrderHandler(*orderUsecase)

	// Create Fiber app
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",                    // Allow frontend
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",              // HTTP methods
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization", // Headers
		AllowCredentials: true,                                       // Allow cookies
	}))

	// Logger middleware
	app.Use(logger.New())

	// Handle preflight requests
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	handlers := &router.Handlers{
		User:       userHandler,
		Coupon:     couponHandler,
		UserCoupon: userCouponHandler,
		Cloth:      clothHandler,
		ClothList:  clothListHandler,
		Order:      orderHandler,
	}

	config.SeedDatabase(db)
	router.SetupRoutes(app, handlers)
	log.Fatal(app.Listen(":18080"))
}
