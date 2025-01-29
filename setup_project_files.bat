@echo off
setlocal enabledelayedexpansion

:: ตรวจสอบว่าอยู่ในโฟลเดอร์ที่ต้องการหรือไม่
echo ✅ Initializing project structure in: %CD%

:: ตรวจสอบและสร้างโฟลเดอร์หากไม่มีอยู่
if not exist cmd mkdir cmd
if not exist config mkdir config
if not exist internal mkdir internal
if not exist internal\domain mkdir internal\domain
if not exist internal\repository mkdir internal\repository
if not exist internal\service mkdir internal\service
if not exist internal\handler mkdir internal\handler
if not exist migrations mkdir migrations
if not exist test mkdir test

:: สร้างไฟล์เปล่าถ้าหากไม่มีอยู่
if not exist cmd\main.go echo. > cmd\main.go
if not exist config\database.go echo. > config\database.go
if not exist internal\domain\models.go echo. > internal\domain\models.go
if not exist internal\repository\emission_repo.go echo. > internal\repository\emission_repo.go
if not exist internal\service\carbon_service.go echo. > internal\service\carbon_service.go
if not exist internal\handler\carbon_handler.go echo. > internal\handler\carbon_handler.go
if not exist migrations\001_create_emission_factors.up.sql echo. > migrations\001_create_emission_factors.up.sql
if not exist migrations\001_create_emission_factors.down.sql echo. > migrations\001_create_emission_factors.down.sql
if not exist test\service_test.go echo. > test\service_test.go

:: ตรวจสอบว่ามี go.mod หรือไม่ ถ้ายังไม่มีให้รัน go mod init
if not exist go.mod (
    for %%I in (.) do set MODULE_NAME=%%~nxI
    go mod init github.com/yourusername/!MODULE_NAME!
)

:: สร้างไฟล์ .gitignore ถ้ายังไม่มี
if not exist .gitignore (
    (
    echo bin/
    echo vendor/
    echo *.exe
    echo *.log
    echo *.out
    echo .DS_Store
    ) > .gitignore
)

:: แสดงผลโครงสร้างไฟล์ที่ถูกสร้าง
echo.
echo ✅ Project files have been created successfully!
echo.
echo 📂 Project Structure:
echo ├── cmd/
echo │   └── main.go
echo ├── config/
echo │   └── database.go
echo ├── internal/
echo │   ├── domain/
echo │   │   └── models.go
echo │   ├── repository/
echo │   │   └── emission_repo.go
echo │   ├── service/
echo │   │   └── carbon_service.go
echo │   ├── handler/
echo │   │   └── carbon_handler.go
echo ├── migrations/
echo │   ├── 001_create_emission_factors.up.sql
echo │   ├── 001_create_emission_factors.down.sql
echo ├── test/
echo │   └── service_test.go
echo ├── .gitignore
echo ├── go.mod
echo.
echo 🚀 Ready to start coding!
