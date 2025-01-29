package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
)

// AssignRequestID Middleware สำหรับเพิ่ม X-Request-ID
func AssignRequestID(c *fiber.Ctx) error {
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
		c.Set("X-Request-ID", requestID)
	}
	c.Locals("request_id", requestID) // ✅ เก็บไว้ใน Fiber Context
	return c.Next()
}

// RequestLogger เพิ่ม Logging ให้ API Requests และแสดง Request ID
func RequestLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${time}] | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${locals:request_id}\n",
	})
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

// GZIPCompression เปิดใช้งานการบีบอัด Response
func GZIPCompression(c *fiber.Ctx) error {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // หรือ compress.LevelBestCompression
	})(c)
}
