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
	clothRepo := postgres.NewClothRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	couponUsecase := usecase.NewCouponUsecase(couponRepo)
	userCouponUsecase := usecase.NewUserCouponUsecase(userCouponRepo)
	clothUsecase := usecase.NewClothUsecase(clothRepo)

	userHandler := handler.NewUserHandler(*userUsecase)
	couponHandler := handler.NewCouponHandler(*couponUsecase)
	userCouponHandler := handler.NewUserCouponHandler(*userCouponUsecase)
	clothHandler := handler.NewClothHandler(*clothUsecase)

	app := fiber.New()

	handlers := &router.Handlers{
		User: userHandler,
		Coupon: couponHandler,
		UserCoupon: userCouponHandler,
		Cloth: clothHandler,
	}
	
	router.SetupRoutes(app, handlers)
	log.Fatal(app.Listen(":3000"))
}