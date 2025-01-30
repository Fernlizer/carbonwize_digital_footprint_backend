package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/config"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/repository"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/routes"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/Fernlizer/carbonwize_digital_footprint_backend/docs"
)

func main() {
	// โหลด Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// เชื่อมต่อ Database
	db := config.InitDB(cfg)

	// สร้าง Repository, Service และ Handler
	carbonRepo := repository.NewEmissionRepository(db)
	carbonService := service.NewCarbonService(carbonRepo)
	carbonHandler := handler.NewCarbonHandler(carbonService)

	// สร้าง Fiber App
	app := fiber.New()

	// ตั้งค่า Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// ตั้งค่า Routes
	routes.SetupRoutes(app, carbonHandler)

	// Graceful Shutdown
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
			log.Fatal("Server failed:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}

	log.Println("Server shutdown successfully!")
}
