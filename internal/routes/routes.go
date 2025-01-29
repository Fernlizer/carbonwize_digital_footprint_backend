package routes

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/config"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/handler"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

// SetupRoutes ตั้งค่า API Routes ทั้งหมด
func SetupRoutes(app *fiber.App, carbonHandler *handler.CarbonHandler) {
	//ใช้ AssignRequestID เพื่อเพิ่ม Request ID ก่อน Log
	app.Use(middleware.AssignRequestID)
	app.Use(middleware.RequestLogger())

	// ใช้ Middleware Health Check ตามมาตรฐาน Kubernetes
	app.Get("/health", healthcheck.New())

	app.Use(healthcheck.New(healthcheck.Config{
		// 🔹 Liveness Probe → ใช้ตรวจสอบว่า Server ยังทำงานอยู่หรือไม่
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true // ถ้า Server ยังรันอยู่ ให้ return true
		},
		LivenessEndpoint: "/live", // เช็ค Liveness ที่ `/live`

		// 🔹 Readiness Probe → ใช้ตรวจสอบว่า Service ทั้งหมดพร้อมทำงานหรือไม่
		ReadinessProbe: func(c *fiber.Ctx) bool {
			db := config.DB
			sqlDB, err := db.DB()
			if err != nil {
				return false
			}
			if err := sqlDB.Ping(); err != nil {
				return false // ถ้า Database ใช้งานไม่ได้ ให้ return false
			}
			return true // ถ้าทุกอย่างพร้อมใช้งาน ให้ return true
		},
		ReadinessEndpoint: "/ready", // เช็ค Readiness ที่ `/ready`
	}))

	api := app.Group("/api/carbon/footprint", middleware.Recover)

	// Apply Global Middleware
	api.Use(middleware.CORS)
	api.Use(middleware.RateLimit)
	api.Use(middleware.GZIPCompression)

	// API Endpoints
	api.Post("/calculate", carbonHandler.CalculateEmissionBasic)
	api.Post("/calculate/weight", carbonHandler.CalculateEmissionWithWeight)
}
