package service

import (
	"errors"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/repository"
)

type CarbonService interface {
	CalculateEmissionBasic(activityType string, distance float64, vehicleType, fuelType string) (float64, error)
	CalculateEmissionWithWeight(activityType string, distance float64, vehicleType, fuelType string, weight float64) (float64, error)
}

type carbonService struct {
	repo repository.EmissionRepository
}

func NewCarbonService(repo repository.EmissionRepository) CarbonService {
	return &carbonService{repo: repo}
}

// ✅ 1. คำนวณแบบพื้นฐาน (ตามโจทย์ ไม่ใช้ `weight`)
func (s *carbonService) CalculateEmissionBasic(activityType string, distance float64, vehicleType, fuelType string) (float64, error) {
	emissionFactor, err := s.repo.GetEmissionFactor(activityType, vehicleType, fuelType)
	if err != nil {
		return 0, err
	}

	// คำนวณตามโจทย์
	totalEmission := distance * emissionFactor
	return totalEmission, nil
}

// ✅ 2. คำนวณโดยใช้ `weight` (สำหรับ ship, airplane ฯลฯ)
func (s *carbonService) CalculateEmissionWithWeight(activityType string, distance float64, vehicleType, fuelType string, weight float64) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("weight is required for ships and airplanes")
	}

	emissionFactor, err := s.repo.GetEmissionFactor(activityType, vehicleType, fuelType)
	if err != nil {
		return 0, err
	}

	// คำนวณโดยใช้ `weight`
	totalEmission := (distance * emissionFactor) * (weight / 1000.0)
	return totalEmission, nil
}
