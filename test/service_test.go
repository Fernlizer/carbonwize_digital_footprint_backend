package test

import (
	"testing"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCalculateEmission(t *testing.T) {
	mockService := service.NewCarbonService(&MockEmissionRepo{})
	result, err := mockService.CalculateEmission("transportation", 100, "car", "gasoline")

	assert.Nil(t, err)
	assert.Equal(t, 25.4, result)
}

type MockEmissionRepo struct{}

func (m *MockEmissionRepo) GetEmissionFactor(activityType, vehicleType, fuelType string) (float64, error) {
	return 0.254, nil
}
