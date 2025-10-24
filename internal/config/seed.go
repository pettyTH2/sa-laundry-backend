package config

import (
	"laundry-backend/internal/entity"
	"gorm.io/gorm"
	"fmt"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64
	db.Model(&entity.User{}).Count(&count)
	if count == 0 {
		users := []entity.User{
			{PhoneNumber: "0811111111", Name: "Jiramet Tangtrairattanakul", Nickname: "Petch", Password: "$2a$10$qoipqMPO15GsAUyKVG0dfOpOXDBqBe39UQTlv7x.Zut1w/tLuXz5W", Role: "customer"},
			{PhoneNumber: "0822222222", Name: "Admin User", Nickname: "Admin", Password: "$2a$10$e5DeRyvOZMSThIRY5QLpcepdlVkDbRflRHUvzcvETlFpapk9F2MBW", Role: "admin"},
			{PhoneNumber: "0833333333", Name: "Laundry User", Nickname: "LaundryAttendant", Password: "$2a$10$G8BRp8o8FXHJDPZ56PzOSeSJynRFlDlS52MEmBJ9078ztLcg/Kvtq", Role: "laundryAttendant"},
			{PhoneNumber: "0844444444", Name: "Cashier User", Nickname: "Cashier", Password: "$2a$10$isl33P5yzwfPcB/YtYxJS.CeiSzuHNe2mKID71aC0Xukxm6lCiWZa", Role: "cashier"},
		}
		for _, user := range users {
			if err := db.Create(&user).Error; err != nil {
				return err
			}
		}
		fmt.Println("Seeded users")
	} else {
		fmt.Println("Users already seeded")
	}

	db.Model(&entity.Coupon{}).Count(&count)
	if count == 0 {
		coupons := []entity.Coupon{
			{CPName: "Machine", CPPrice: 900},
			{CPName: "Hand Wash", CPPrice: 1250},
			{CPName: "Iron", CPPrice: 750},
		}
		for _, coupon := range coupons {
			if err := db.Create(&coupon).Error; err != nil {
				return err
			}
		}
		fmt.Println("Seeded coupons")
	} else {
		fmt.Println("Coupons already seeded")
	}

	db.Model(&entity.Cloth{}).Count(&count)
	if count == 0 {
		clothes := []entity.Cloth{
			{ClothName: "Long-Shirt", ClothPrice: 25, Category: "Wash Dry"},
			{ClothName: "Short-Shirt", ClothPrice: 20, Category: "Wash Dry"},
			{ClothName: "Short", ClothPrice: 20, Category: "Wash Dry"},
			{ClothName: "Blouse", ClothPrice: 25, Category: "Wash Dry"},
			{ClothName: "Dress", ClothPrice: 80, Category: "Wash Dry"},
			{ClothName: "Suit", ClothPrice: 150, Category: "Dry Clean"},
			{ClothName: "Safari Suit", ClothPrice: 150, Category: "Dry Clean"},
			{ClothName: "Sweater", ClothPrice: 250, Category: "Dry Clean"},
			{ClothName: "Jacket", ClothPrice: 120, Category: "Dry Clean"},
			{ClothName: "Dress", ClothPrice: 150, Category: "Dry Clean"},
		}
		for _, cloth := range clothes {
			if err := db.Create(&cloth).Error; err != nil {
				return err
			}
		}
		fmt.Println("Seeded clothes")
	} else {
		fmt.Println("Clothes already seeded")
	}

	db.Model(&entity.UserCoupon{}).Count(&count)
	if count == 0 {
		userCoupons := []entity.UserCoupon{
			{PointLeft: 32, StartDate: "01/01/2000", ExpireDate: "01/02/2000", UserID: 1, CouponID: 1},
			{PointLeft: 50, StartDate: "02/03/2000", ExpireDate: "02/04/2000", UserID: 1, CouponID: 2},
		}
		for _, userCoupon := range userCoupons {
			if err := db.Create(&userCoupon).Error; err != nil {
				return err
			}
		}
		fmt.Println("Seeded userCoupon")
	} else {
		fmt.Println("UserCoupon already seeded")
	}

	db.Model(&entity.Order{}).Count(&count)
	if count == 0 {
		orders := []entity.Order{
			{ServiceType: "Machine", TotalCloth: 18, TotalCost: 18, OrderDate: "01/01/2000", PickupDate: "04/01/2000", OrderStatus: "Completed", PaymentMethod: "Package", UserID: 1},
			{ServiceType: "Dry Clean", TotalCloth: 5, TotalCost: 750, OrderDate: "05/01/2000", PickupDate: "08/01/2000", OrderStatus: "Pending", PaymentMethod: "Cash", UserID: 1},
		}
		for _, order := range orders {
			if err := db.Create(&order).Error; err != nil {
				return err
			}
		}
		fmt.Println("Seeded order")
	} else {
		fmt.Println("Order already seeded")
	}

	db.Model(&entity.ClothList{}).Count(&count)
	if count == 0 {
		clothLists := []entity.ClothList{
			{Quantity: 15, SubTotalCost: 15, OrderID: 1, ClothID: 1},
			{Quantity: 3, SubTotalCost: 3, OrderID: 1, ClothID: 3},
			{Quantity: 5, SubTotalCost: 750, OrderID: 2, ClothID: 6},
		}
		for _, clostList := range clothLists {
			if err := db.Create(&clostList).Error; err != nil {
				return err
			}
		}
		fmt.Println("Seeded clothtList")
	} else {
		fmt.Println("ClothtList already seeded")
	}

	return nil
}