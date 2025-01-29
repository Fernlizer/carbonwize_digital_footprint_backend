package handler

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/domain"
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

// CarbonHandler จัดการ API Carbon Footprint
type CarbonHandler struct {
	service service.CarbonService
}

func NewCarbonHandler(s service.CarbonService) *CarbonHandler {
	return &CarbonHandler{service: s}
}

// CalculateEmissionBasic คำนวณ Carbon Footprint แบบปกติ
// @Summary คำนวณ Carbon Footprint
// @Description รับค่ากิจกรรมและระยะทาง แล้วคืนค่า Carbon Footprint
// @Tags Carbon
// @Accept json
// @Produce json
// @Param request body domain.CarbonRequest true "Request Body"
// @Success 200 {object} domain.CarbonResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /carbon/footprint/calculate [post]
func (h *CarbonHandler) CalculateEmissionBasic(c *fiber.Ctx) error {
	var req domain.CarbonRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}

	emission, err := h.service.CalculateEmissionBasic(req.ActivityType, req.DistanceKm, req.VehicleType, req.FuelType)
	if err != nil {
		return c.Status(500).JSON(domain.ErrorResponse{Error: "Calculation failed"})
	}

	return c.JSON(domain.CarbonResponse{
		ActivityType:     req.ActivityType,
		CarbonEmissionKg: emission,
	})
}

// CalculateEmissionWithWeight คำนวณ Carbon Footprint แบบใช้ Weight
// @Summary คำนวณ Carbon Footprint (พาหนะที่มีน้ำหนัก)
// @Description รับค่ากิจกรรม, ระยะทาง และน้ำหนัก แล้วคืนค่า Carbon Footprint
// @Tags Carbon
// @Accept json
// @Produce json
// @Param request body domain.CarbonRequestWithWeight true "Request Body"
// @Success 200 {object} domain.CarbonResponseWithWeight
// @Failure 400 {object} domain.ErrorResponse
// @Router /carbon/footprint/calculate/weight [post]
func (h *CarbonHandler) CalculateEmissionWithWeight(c *fiber.Ctx) error {
	var req domain.CarbonRequestWithWeight
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}

	emission, err := h.service.CalculateEmissionWithWeight(req.ActivityType, req.DistanceKm, req.VehicleType, req.FuelType, req.Weight)
	if err != nil {
		return c.Status(500).JSON(domain.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(domain.CarbonResponseWithWeight{
		ActivityType:     req.ActivityType,
		CarbonEmissionKg: emission,
		WeightTons:       req.Weight / 1000.0,
	})
}
