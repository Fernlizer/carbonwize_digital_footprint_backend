package domain

import "gorm.io/gorm"

type EmissionFactor struct {
	gorm.Model
	ActivityType   string  `gorm:"type:varchar(50);not null"`
	VehicleType    string  `gorm:"type:varchar(50);not null"`
	FuelType       string  `gorm:"type:varchar(50);not null"`
	EmissionFactor float64 `gorm:"type:numeric;not null"`
}
