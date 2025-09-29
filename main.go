package main

import (
	"laundry-backend/internal/database"
	"laundry-backend/internal/handler"
	"laundry-backend/internal/repository/postgres"
	"laundry-backend/internal/usecase"
	"laundry-backend/internal/router"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.NewDatabaseConnection()

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}

	userRepo := postgres.NewUserRepository(db)
	couponRepo := postgres.NewCouponRepository(db)
	userCouponRepo := postgres.NewUserCouponRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	couponUsecase := usecase.NewCouponUsecase(couponRepo)
	userCouponUsecase := usecase.NewUserCouponUsecase(userCouponRepo)

	userHandler := handler.NewUserHandler(*userUsecase)
	couponHandler := handler.NewCouponHandler(*couponUsecase)
	userCouponHandler := handler.NewUserCouponHandler(*userCouponUsecase)

	app := fiber.New()

	handlers := &router.Handlers{
		User: userHandler,
		Coupon: couponHandler,
		UserCoupon: userCouponHandler,

	}
	
	router.SetupRoutes(app, handlers)
	log.Fatal(app.Listen(":3000"))
}