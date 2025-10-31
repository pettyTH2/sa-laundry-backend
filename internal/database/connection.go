package database

import (
	"fmt"
	"log"
	"laundry-backend/internal/config"
	"laundry-backend/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func NewDatabaseConnection() *gorm.DB {
	cfg := config.LoadDBConfig()
	dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
    )
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
    	log.Fatal("Failed to connect to database:", err)
  	}
	return db
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
        &entity.Order{},      
        &entity.User{},        
        &entity.Coupon{},       
        &entity.Cloth{},        
        &entity.UserCoupon{}, 
        &entity.ClothList{},
    )
}