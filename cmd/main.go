package main

import (
	"fmt"
	"log"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/config"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/docs"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/repository"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/routes"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// โหลด Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// เชื่อมต่อ Database
	db := config.InitDB(cfg)

	// สร้าง Repository, Service และ Handler
	carbonRepo := repository.NewEmissionRepository(db)
	carbonService := service.NewCarbonService(carbonRepo)
	carbonHandler := handler.NewCarbonHandler(carbonService)

	// สร้าง Fiber App
	app := fiber.New()

	// ตั้งค่า Swagger
	docs.SwaggerInfo.Title = "CarbonWize API"
	docs.SwaggerInfo.Description = "API สำหรับคำนวณ Carbon Footprint"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + cfg.AppPort
	docs.SwaggerInfo.BasePath = "/api"

	// เพิ่ม Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// ตั้งค่า Routes
	routes.SetupRoutes(app, carbonHandler)

	// รันเซิร์ฟเวอร์
	port := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("🚀 Server is running on http://localhost%s", port)
	log.Fatal(app.Listen(port))
}
