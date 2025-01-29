package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// RequestLogger เพิ่ม Logging ให้ API Requests
func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	log.Printf("🔹 %s %s | %v | %d", c.Method(), c.OriginalURL(), duration, c.Response().StatusCode())
	return err
}

// Recover จัดการ Panic และคืนค่า JSON แทนการ Crash
func Recover(c *fiber.Ctx) error {
	return recover.New()(c)
}

// CORS เปิดให้ API รองรับ Cross-Origin Requests
func CORS(c *fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
		MaxAge:           86400,
	})(c)
}

// RateLimit ป้องกันการยิง API ถี่เกินไป
func RateLimit(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        100,             // 100 requests
		Expiration: 1 * time.Minute, // ต่อ 1 นาที
	})(c)
}
