package handler

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CarbonHandler struct {
	service service.CarbonService
}

func NewCarbonHandler(s service.CarbonService) *CarbonHandler {
	return &CarbonHandler{service: s}
}

func (h *CarbonHandler) CalculateEmission(c *fiber.Ctx) error {
	type Request struct {
		ActivityType string  `json:"activity_type"`
		DistanceKm   float64 `json:"distance_km"`
		VehicleType  string  `json:"vehicle_type"`
		FuelType     string  `json:"fuel_type"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	emission, err := h.service.CalculateEmission(req.ActivityType, req.DistanceKm, req.VehicleType, req.FuelType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Calculation failed"})
	}

	return c.JSON(fiber.Map{
		"activity_type":      req.ActivityType,
		"carbon_emission_kg": emission,
	})
}
