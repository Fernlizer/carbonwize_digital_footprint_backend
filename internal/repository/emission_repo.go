package repository

import (
	"github.com/Fernlizer/carbonwize_digital_footprint_backend/internal/domain"
	"gorm.io/gorm"
)

type EmissionRepository interface {
	GetEmissionFactor(activityType, vehicleType, fuelType string) (float64, error)
}

type emissionRepo struct {
	db *gorm.DB
}

func NewEmissionRepository(db *gorm.DB) EmissionRepository {
	return &emissionRepo{db: db}
}

func (r *emissionRepo) GetEmissionFactor(activityType, vehicleType, fuelType string) (float64, error) {
	var emission domain.EmissionFactor
	err := r.db.Where("activity_type = ? AND vehicle_type = ? AND fuel_type = ?", activityType, vehicleType, fuelType).First(&emission).Error
	if err != nil {
		return 0, err
	}
	return emission.EmissionFactor, nil
}
