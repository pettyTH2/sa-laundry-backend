package database

import (
	"fmt"
	"log"
	"laundry-backend/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
  host     = "localhost"  
  port     = 5432         
  user     = "laundry_user"     
  password = "mypassword" 
  dbname   = "laundry_db" 
)

func NewDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
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