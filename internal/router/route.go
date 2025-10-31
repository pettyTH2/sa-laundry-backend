package router

import (
	"laundry-backend/internal/handler"
	"laundry-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	User       *handler.UserHandler
	Coupon     *handler.CouponHandler
	UserCoupon *handler.UserCouponHandler
	Cloth      *handler.ClothHandler
	ClothList  *handler.ClothListHandler
	Order      *handler.OrderHandler
}

func SetupRoutes(app *fiber.App, h *Handlers) {

	// Public routes (no auth required)
	app.Post("/users/login", h.User.UserLogin)

	// Protected routes (auth required)
	app.Use(middleware.RequireAuth)

	// User routes
	users := app.Group("/users")
	users.Post("/register", middleware.RequireStaffAuth, h.User.CreateUser)
	users.Get("/role/:role", h.User.GetUserByRole)
	users.Get("/id/:id", h.User.GetUserById)
	users.Get("/", h.User.GetAllUsers)
	users.Put("/:id", h.User.UpdateUser)

	coupons := app.Group("/coupons")
	coupons.Post("/", h.Coupon.CreateCoupon)
	coupons.Get("/:id", h.Coupon.GetCouponByID)
	coupons.Get("/", h.Coupon.GetAllCoupons)
	coupons.Put("/:id", h.Coupon.UpdateCoupon)

	userCoupons := app.Group("/user_coupons")
	userCoupons.Post("/", h.UserCoupon.CreateUserCoupon)
	userCoupons.Get("/user/:user_id", h.UserCoupon.GetUserCouponsByUserID)
	userCoupons.Get("/coupon/:coupon_id", h.UserCoupon.GetUserCouponsByCouponID)
	userCoupons.Get("/", h.UserCoupon.GetAllUserCoupons)
	userCoupons.Put("/:id", h.UserCoupon.UpdateUserCoupon)

	cloths := app.Group("/cloths")
	cloths.Post("/", h.Cloth.CreateCloth)
	cloths.Get("/category/:category", h.Cloth.GetClothsByCategory)
	cloths.Get("/", h.Cloth.GetAllCloths)
	cloths.Put("/:id", h.Cloth.UpdateCloth)

	clothLists := app.Group("/cloth_lists")
	clothLists.Post("/", h.ClothList.CreateClothList)
	clothLists.Get("/order/:order_id", h.ClothList.GetClothListsByOrderID)
	clothLists.Put("/:id", h.ClothList.UpdateClothList)

	orders := app.Group("/orders")
	orders.Post("/", h.Order.CreateOrder)
	orders.Get("/user/:user_id", h.Order.GetOrdersByUserID)
	orders.Get("/", h.Order.GetAllOrders)
	orders.Put("/:id", h.Order.UpdateOrder)
	orders.Get("/:id", h.Order.GetOrderByID)
}
