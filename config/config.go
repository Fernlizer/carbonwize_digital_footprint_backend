package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	AppPort      string `mapstructure:"APP_PORT"`
	DatabaseHost string `mapstructure:"DB_HOST"`
	DatabaseUser string `mapstructure:"DB_USER"`
	DatabasePass string `mapstructure:"DB_PASSWORD"`
	DatabaseName string `mapstructure:"DB_NAME"`
	DatabasePort string `mapstructure:"DB_PORT"`
}

// LoadConfig ใช้สำหรับโหลดค่าการตั้งค่าจากไฟล์และ Environment Variables
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // ใช้ไฟล์ชื่อ config
	viper.SetConfigType("yaml")   // ใช้ไฟล์ .yaml
	viper.AddConfigPath(".")      // ค้นหาไฟล์ในโฟลเดอร์ปัจจุบัน

	// อ่านค่าจาก Environment Variables ด้วย
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
		return nil, err
	}

	fmt.Println("✅ Config loaded successfully!")
	return &config, nil
}
