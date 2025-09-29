package router
import (
	"laundry-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	User *handler.UserHandler
	Coupon *handler.CouponHandler
	UserCoupon *handler.UserCouponHandler
}

func SetupRoutes(app *fiber.App, h *Handlers) {

	users := app.Group("/users")
	users.Post("/", h.User.CreateUser)
	users.Get("/role/:role", h.User.GetUserByRole)
	users.Get("/phone/:phone_number", h.User.GetUserByPhoneNumber)
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
}