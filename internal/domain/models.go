package domain

import "gorm.io/gorm"

type EmissionFactor struct {
	gorm.Model
	ActivityType   string  `gorm:"type:varchar(50);not null"`
	VehicleType    string  `gorm:"type:varchar(50);not null"`
	FuelType       string  `gorm:"type:varchar(50);not null"`
	EmissionFactor float64 `gorm:"type:numeric;not null"`
}

// CarbonRequest ใช้สำหรับ API Basic
type CarbonRequest struct {
	ActivityType string  `json:"activity_type" example:"transportation"`
	DistanceKm   float64 `json:"distance_km" example:"100"`
	VehicleType  string  `json:"vehicle_type" example:"car"`
	FuelType     string  `json:"fuel_type" example:"gasoline"`
}

// CarbonRequestWithWeight ใช้สำหรับ API Advanced
type CarbonRequestWithWeight struct {
	ActivityType string  `json:"activity_type" example:"transportation"`
	DistanceKm   float64 `json:"distance_km" example:"1000"`
	VehicleType  string  `json:"vehicle_type" example:"airplane"`
	FuelType     string  `json:"fuel_type" example:"jet_fuel"`
	Weight       float64 `json:"weight" example:"80000"`
}

// CarbonResponse ใช้สำหรับ API Basic
type CarbonResponse struct {
	ActivityType     string  `json:"activity_type" example:"transportation"`
	CarbonEmissionKg float64 `json:"carbon_emission_kg" example:"25.4"`
}

// CarbonResponseWithWeight ใช้สำหรับ API Advanced
type CarbonResponseWithWeight struct {
	ActivityType     string  `json:"activity_type" example:"transportation"`
	CarbonEmissionKg float64 `json:"carbon_emission_kg" example:"20000"`
	WeightTons       float64 `json:"weight_tons" example:"80"`
}

// ErrorResponse ใช้สำหรับ Error Handling
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid input"`
}
