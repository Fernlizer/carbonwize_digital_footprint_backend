# CarbonWize Digital Footprint Backend 🚀

🔥 **พร้อมใช้งานระดับ Production!** 🔥

✅ **CI/CD พร้อมใช้งาน** → ทุกการ push และ pull request บน `main` และ `DEV` จะ **รัน Unit Tests อัตโนมัติ** เพื่อให้แน่ใจว่าโค้ดมีคุณภาพสูงสุด!
✅ **Swagger API Documentation** → รองรับการใช้งานและทดลอง API ผ่าน **Swagger UI**
✅ **Unit Testing & Code Coverage** → **มั่นใจได้ว่าโค้ดทำงานถูกต้องทุกฟังก์ชัน** ด้วยชุดทดสอบที่ครอบคลุม
✅ **Database Migration & Seed Data** → ใช้ **Golang + PostgreSQL** พร้อมระบบ **Migration** และ **ค่าเริ่มต้น**
✅ **Fiber Framework** → **เร็ว แรง ทันสมัย** รองรับ **Middleware, Routing และ Error Handling** ที่มีประสิทธิภาพ

---

## 🔹 **1. Setup Project**
### ✅ **1.1 ติดตั้ง Dependency**
ให้รันคำสั่งนี้เพื่อโหลด package ทั้งหมดที่จำเป็น
```sh
go mod tidy
```

---

### ✅ **1.2 ตั้งค่า Database**
📌 **ใช้ PostgreSQL และให้สร้าง Database ชื่อ `carbon_db`**  
📌 **ตั้งค่าไฟล์ `config.yaml`**  
```yaml
APP_PORT: "8080"
DB_HOST: "localhost"
DB_USER: "postgres"
DB_PASSWORD: "yourpassword"
DB_NAME: "carbon_db"
DB_PORT: "5432"
```
---

### ✅ **1.3 รัน Migration**
📌 **สร้างตารางและเพิ่มข้อมูลเริ่มต้น**
```sh
migrate -database "postgres://postgres:yourpassword@localhost:5432/carbon_db?sslmode=disable" -path migrations up
```

---

### ✅ **1.4 รันโปรเจกต์**
📌 **ใช้ `Air` เพื่อรันแบบ Hot Reload**
```sh
air
```
📌 **หรือใช้ `go run`**
```sh
go run cmd/main.go
```

---

### ✅ **1.5 ตั้งค่า Swagger API Documentation**
📌 **สร้างไฟล์ Swagger Docs**
```sh
swag init -g cmd/main.go --output ./docs
```
📌 **รันโปรเจกต์ และเปิด Swagger UI**
```sh
go run cmd/main.go
```
👉 **เปิด Swagger UI ได้ที่:** `http://localhost:8080/swagger/index.html`

---

## 🔹 **2. CI/CD Pipeline - Auto Testing ✅**
ทุกครั้งที่มีการ **push** หรือ **pull request** ไปยัง `main` และ `DEV` ระบบจะทำการ **รัน Unit Tests อัตโนมัติ** บน GitHub Actions

📌 **ตัวอย่าง Workflow:** `.github/workflows/ci.yml`
```yaml
name: CI Pipeline - Auto Testing

on:
  push:
    branches:
      - main
      - DEV
  pull_request:
    branches:
      - main
      - DEV

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
      - name: 📥 Checkout code
        uses: actions/checkout@v3

      - name: 🔧 Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: '1.23'

      - name: 🏗 Verify Go version
        run: go version

      - name: 📦 Install dependencies
        run: |
          go env  # ✅ ตรวจสอบ GOPATH และ Environment Variables
          go mod tidy

      - name: ✅ Run Unit Tests
        run: go test -v ./test

      - name: 📊 Check test coverage
        run: go test -cover ./test
```
📌 **ไปที่ GitHub → แท็บ `Actions` → ดูผลลัพธ์ของ CI/CD**

---

## 🔹 **3. ทดสอบ API ด้วย Postman หรือ Curl**
📌 **คำนวณ Carbon Footprint (แบบไม่มีน้ำหนัก)**  
```sh
curl -X 'POST' \
  'http://localhost:8080/api/carbon/footprint/calculate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "activity_type": "transportation",
  "distance_km": 100,
  "vehicle_type": "car",
  "fuel_type": "gasoline"
}'
```
📌 **คำนวณ Carbon Footprint (แบบใช้ `weight` เช่น เครื่องบิน / เรือ)**  
```sh
curl -X 'POST' \
  'http://localhost:8080/api/carbon/footprint/calculate/weight' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "activity_type": "transportation",
  "distance_km": 1000,
  "vehicle_type": "airplane",
  "fuel_type": "jet_fuel",
  "weight": 80000
}'
```

---

## 🎯 **สรุป**
| ✅ Feature | 📌 รายละเอียด |
|-------------|----------------|
| **🔥 Ready for Production** | **โครงสร้างพร้อมใช้งานจริง!** 🚀 |
| **✅ CI/CD Pipeline** | **Auto Testing ทุกครั้งที่ Push หรือ PR!** 📊 |
| **📜 Swagger API Docs** | **Swagger UI พร้อมให้ใช้งาน!** 📖 |
| **🛠️ Unit Tests & Coverage** | **มั่นใจได้ว่าโค้ดทำงานถูกต้อง!** ✅ |
| **🐳 ใช้ Fiber Framework** | **เร็ว แรง พัฒนา Products ได้ไว** ⚡ |

