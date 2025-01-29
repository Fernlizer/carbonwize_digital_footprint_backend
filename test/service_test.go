package test

import (
	"testing"

	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/service"
	"github.com/stretchr/testify/assert"
)

// ✅ ทดสอบ CalculateEmissionBasic (ไม่ใช้ weight)
func TestCalculateEmissionBasic(t *testing.T) {
	mockService := service.NewCarbonService(&MockEmissionRepo{})
	result, err := mockService.CalculateEmissionBasic("transportation", 100, "car", "gasoline")

	assert.Nil(t, err)
	assert.Equal(t, 25.4, result)
}

// ✅ ทดสอบ CalculateEmissionWithWeight (สำหรับเครื่องบินและเรือ)
func TestCalculateEmissionWithWeight(t *testing.T) {
	mockService := service.NewCarbonService(&MockEmissionRepo{})
	result, err := mockService.CalculateEmissionWithWeight("transportation", 1000, "airplane", "jet_fuel", 80000)

	assert.Nil(t, err)
	assert.Equal(t, 20000.0, result) // 1000 km * 0.25 * (80000 / 1000)
}

// ✅ Mock Repository
type MockEmissionRepo struct{}

// ✅ Mock ค่า emission factor ตามที่กำหนด
func (m *MockEmissionRepo) GetEmissionFactor(activityType, vehicleType, fuelType string) (float64, error) {
	mockData := map[string]float64{
		"transportation:car:gasoline":      0.254,
		"transportation:airplane:jet_fuel": 0.25,
	}

	key := activityType + ":" + vehicleType + ":" + fuelType
	if val, ok := mockData[key]; ok {
		return val, nil
	}
	return 0, nil
}
