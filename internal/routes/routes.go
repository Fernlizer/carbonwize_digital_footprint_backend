package routes

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes ตั้งค่า API Routes ทั้งหมด
func SetupRoutes(app *fiber.App, carbonHandler *handler.CarbonHandler) {
	// Group API Routes
	api := app.Group("/api/carbon/footprint", middleware.RequestLogger, middleware.Recover)

	// Apply Global Middleware (CORS, Rate Limiting, Auth)
	api.Use(middleware.CORS)
	api.Use(middleware.RateLimit)

	// API Endpoints
	api.Post("/calculate", carbonHandler.CalculateEmissionBasic)
	api.Post("/calculate/weight", carbonHandler.CalculateEmissionWithWeight)
}
