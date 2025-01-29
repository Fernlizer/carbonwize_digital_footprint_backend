package main

import (
	"fmt"
	"log"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/config"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/repository"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/routes"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// โหลด Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// เชื่อมต่อ Database (ถ้าไม่มีตารางหรือเชื่อมต่อไม่ได้ -> หยุดโปรแกรม)
	db := config.InitDB(cfg)

	// สร้าง Repository, Service และ Handler
	carbonRepo := repository.NewEmissionRepository(db)
	carbonService := service.NewCarbonService(carbonRepo)
	carbonHandler := handler.NewCarbonHandler(carbonService)

	// ตั้งค่า Fiber App
	app := fiber.New()

	// ตั้งค่า Routes
	routes.SetupRoutes(app, carbonHandler)

	// รันเซิร์ฟเวอร์ที่ Port ตาม Config
	port := fmt.Sprintf(":%s", cfg.AppPort)
	log.Fatal(app.Listen(port))
}
