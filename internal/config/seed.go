package config

import (
	"fmt"
	"laundry-backend/internal/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64

	// Seed Users
	db.Model(&entity.User{}).Count(&count)
	if count == 0 {
		users := []entity.User{
			{ID: uuid.New(), PhoneNumber: "0811111111", Name: "Jiramet Tangtrairattanakul", Nickname: "Petch", Password: "$2a$10$qoipqMPO15GsAUyKVG0dfOpOXDBqBe39UQTlv7x.Zut1w/tLuXz5W", Role: "customer"},
			{ID: uuid.New(), PhoneNumber: "0822222222", Name: "Admin User", Nickname: "Admin", Password: "$2a$10$e5DeRyvOZMSThIRY5QLpcepdlVkDbRflRHUvzcvETlFpapk9F2MBW", Role: "admin"},
			{ID: uuid.New(), PhoneNumber: "0833333333", Name: "Laundry User", Nickname: "LaundryAttendant", Password: "$2a$10$G8BRp8o8FXHJDPZ56PzOSeSJynRFlDlS52MEmBJ9078ztLcg/Kvtq", Role: "laundryAttendant"},
			{ID: uuid.New(), PhoneNumber: "0844444444", Name: "Cashier User", Nickname: "Cashier", Password: "$2a$10$isl33P5yzwfPcB/YtYxJS.CeiSzuHNe2mKID71aC0Xukxm6lCiWZa", Role: "cashier"},
		}
		db.Create(&users)
		fmt.Println("Seeded users")

		// Seed Coupons
		coupons := []entity.Coupon{
			{ID: uuid.New(), CPName: "Machine", CPPrice: 900},
			{ID: uuid.New(), CPName: "Hand Wash", CPPrice: 1250},
			{ID: uuid.New(), CPName: "Iron", CPPrice: 750},
		}
		db.Create(&coupons)
		fmt.Println("Seeded coupons")

		// Seed Clothes
		clothes := []entity.Cloth{
			{ID: uuid.New(), ClothName: "Long-Shirt", ClothPrice: 25, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Short-Shirt", ClothPrice: 20, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Long T-Shirt", ClothPrice: 25, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Short T-Shirt", ClothPrice: 20, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Pant", ClothPrice: 25, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Short", ClothPrice: 20, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Jean", ClothPrice: 25, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Long Skirt", ClothPrice: 25, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Short Skirt", ClothPrice: 20, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Blouse", ClothPrice: 25, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Long Blouse", ClothPrice: 30, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Dress", ClothPrice: 80, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Towel", ClothPrice: 30, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Pillow Case", ClothPrice: 15, Category: "Wash Dry"},
			{ID: uuid.New(), ClothName: "Suit", ClothPrice: 150, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Safari Suit", ClothPrice: 150, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Sweater", ClothPrice: 250, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Jacket", ClothPrice: 120, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Dress", ClothPrice: 150, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Shirt", ClothPrice: 80, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Blouse", ClothPrice: 80, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Skirt", ClothPrice: 80, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Pant", ClothPrice: 80, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Scarf", ClothPrice: 60, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Necktie", ClothPrice: 50, Category: "Dry Clean"},
			{ID: uuid.New(), ClothName: "Coat", ClothPrice: 250, Category: "Dry Clean"},
		}
		db.Create(&clothes)
		fmt.Println("Seeded clothes")

		// Seed UserCoupons
		userCoupons := []entity.UserCoupon{
			{ID: uuid.New(), PointLeft: 32, StartDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local), ExpireDate: time.Date(2026, time.February, 1, 0, 0, 0, 0, time.Local), UserID: users[0].ID, CouponID: coupons[0].ID},
			{ID: uuid.New(), PointLeft: 50, StartDate: time.Date(2024, time.March, 2, 0, 0, 0, 0, time.Local), ExpireDate: time.Date(2026, time.April, 2, 0, 0, 0, 0, time.Local), UserID: users[0].ID, CouponID: coupons[1].ID},
		}
		db.Create(&userCoupons)
		fmt.Println("Seeded userCoupon")

		// Seed Orders
		orders := []entity.Order{
			{ID: uuid.New(), ServiceType: "Machine", TotalCloth: 18, TotalCost: 450, OrderDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local), PickupDate: time.Date(2024, time.January, 4, 0, 0, 0, 0, time.Local), OrderStatus: "คำสั่งซื้อเสร็จสิ้น", PaymentMethod: "Coupon", UserID: users[0].ID},
			{ID: uuid.New(), ServiceType: "Dry Clean", TotalCloth: 5, TotalCost: 750, OrderDate: time.Date(2024, time.January, 5, 0, 0, 0, 0, time.Local), PickupDate: time.Date(2024, time.January, 8, 0, 0, 0, 0, time.Local), OrderStatus: "รอดำเนินการ", PaymentMethod: "Cash", UserID: users[0].ID},
		}
		db.Create(&orders)
		fmt.Println("Seeded order")

		// Seed ClothLists
		clothLists := []entity.ClothList{
			{ID: uuid.New(), Quantity: 15, SubTotalCost: 375, OrderID: orders[0].ID, ClothID: clothes[0].ID},
			{ID: uuid.New(), Quantity: 3, SubTotalCost: 75, OrderID: orders[0].ID, ClothID: clothes[2].ID},
			{ID: uuid.New(), Quantity: 5, SubTotalCost: 750, OrderID: orders[1].ID, ClothID: clothes[14].ID},
		}
		db.Create(&clothLists)
		fmt.Println("Seeded clothList")
	} else {
		fmt.Println("Database already seeded")
	}

	return nil
}
