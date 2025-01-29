package service

import "github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/repository"

type CarbonService interface {
	CalculateEmission(activityType string, distance float64, vehicleType, fuelType string) (float64, error)
}

type carbonService struct {
	repo repository.EmissionRepository
}

func NewCarbonService(repo repository.EmissionRepository) CarbonService {
	return &carbonService{repo: repo}
}

func (s *carbonService) CalculateEmission(activityType string, distance float64, vehicleType, fuelType string) (float64, error) {
	emissionFactor, err := s.repo.GetEmissionFactor(activityType, vehicleType, fuelType)
	if err != nil {
		return 0, err
	}
	return distance * emissionFactor, nil
}
