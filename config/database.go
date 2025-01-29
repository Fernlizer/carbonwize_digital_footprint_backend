package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ตัวแปร global ที่ใช้เก็บการเชื่อมต่อ Database
var DB *gorm.DB

// InitDB ใช้สำหรับเชื่อมต่อ PostgreSQL โดยใช้ค่าจาก Config
func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePass, cfg.DatabaseName, cfg.DatabasePort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("✅ Database connected successfully!")
	return DB
}
