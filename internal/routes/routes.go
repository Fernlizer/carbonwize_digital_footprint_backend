package routes

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes ตั้งค่า API Routes ทั้งหมด
func SetupRoutes(app *fiber.App, carbonHandler *handler.CarbonHandler) {
	//ใช้ AssignRequestID เพื่อเพิ่ม Request ID ก่อน Log
	app.Use(middleware.AssignRequestID)
	app.Use(middleware.RequestLogger())

	//Health Check API
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "OK", "message": "CarbonWize API is running"})
	})

	api := app.Group("/api/carbon/footprint", middleware.Recover)

	// Apply Global Middleware
	api.Use(middleware.CORS)
	api.Use(middleware.RateLimit)
	api.Use(middleware.GZIPCompression)

	// API Endpoints
	api.Post("/calculate", carbonHandler.CalculateEmissionBasic)
	api.Post("/calculate/weight", carbonHandler.CalculateEmissionWithWeight)
}
