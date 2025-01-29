package routes

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes ตั้งค่า API Routes ทั้งหมด
func SetupRoutes(app *fiber.App, carbonHandler *handler.CarbonHandler) {
	api := app.Group("/api/carbon/footprint")
	api.Post("/calculate", carbonHandler.CalculateEmission)
}
