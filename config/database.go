package config

import (
	"fmt"
	"log"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/domain"
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

	// เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// ตรวจสอบว่าฐานข้อมูลมีอยู่หรือไม่
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("database is not available. Please check your database connection.")
	}

	// ตรวจสอบว่าตาราง `emission_factors` มีอยู่หรือไม่
	if !db.Migrator().HasTable(&domain.EmissionFactor{}) {
		log.Fatal("table 'emission_factors' does not exist. Please run database migration.")
	}

	// ตั้งค่าให้ใช้งานฐานข้อมูลนี้
	DB = db
	fmt.Println("database connected successfully!")
	return DB
}
