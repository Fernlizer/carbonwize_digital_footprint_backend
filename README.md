# CarbonWize Digital Footprint Backend

**ระบบ Backend สำหรับคำนวณ Carbon Footprint โดยใช้ Golang (Fiber) พร้อม CI/CD และ API Documentation**

## 📌 คุณสมบัติหลัก
- **CI/CD Pipeline** → รัน Unit Tests อัตโนมัติทุกครั้งที่ push หรือ pull request ไปยัง `main` และ `DEV`
- **Swagger API Documentation** → รองรับการใช้งานและทดลอง API ผ่าน **Swagger UI**
- **Unit Testing & Code Coverage** → ทดสอบฟังก์ชันการทำงานของ API ด้วยชุดทดสอบที่ครอบคลุม
- **Database Migration & Seed Data** → ใช้ **PostgreSQL** พร้อมระบบ **Migration** และค่าเริ่มต้น
- **Fiber Framework** → รองรับ Middleware, Routing, และ Error Handling
- **Request ID & Logging** → ใช้ **X-Request-ID** เพื่อติดตาม Request แต่ละรายการ
- **Health Check API** → รองรับ `/live`, `/ready` ตามมาตรฐาน Kubernetes

---

## 📌 **1. ติดตั้งและตั้งค่าโปรเจกต์**

### ✅ **1.1 ติดตั้ง Dependency**
```sh
go mod tidy
```

### ✅ **1.2 ตั้งค่า Database**
ใช้ PostgreSQL และสร้าง Database `carbon_db` จากนั้นตั้งค่าไฟล์ `config.yaml`
```yaml
APP_PORT: "8080"
DB_HOST: "localhost"
DB_USER: "postgres"
DB_PASSWORD: "yourpassword"
DB_NAME: "carbon_db"
DB_PORT: "5432"
```

### ✅ **1.3 รัน Migration**
```sh
migrate -database "postgres://postgres:yourpassword@localhost:5432/carbon_db?sslmode=disable" -path migrations up
```

### ✅ **1.4 รันโปรเจกต์**
```sh
air  # หรือใช้ go run
```
```sh
go run cmd/main.go
```

### ✅ **1.5 สร้างและเรียกใช้ Swagger API Docs**
```sh
swag init -g cmd/main.go --output ./docs
go run cmd/main.go
```
เปิด **Swagger UI** ที่: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 📌 **2. CI/CD Pipeline - Auto Testing**
ทุกครั้งที่มี **push** หรือ **pull request** ไปยัง `main` หรือ `DEV` ระบบจะ **รัน Unit Tests อัตโนมัติ** บน GitHub Actions

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
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: '1.23'

      - name: Install dependencies
        run: |
          go mod tidy

      - name: Run Unit Tests
        run: go test -v ./test

      - name: Check test coverage
        run: go test -cover ./test
```

---

## 📌 **3. ทดสอบ API ด้วย Postman หรือ Curl**
### **คำนวณ Carbon Footprint (แบบไม่มีน้ำหนัก)**
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
### **คำนวณ Carbon Footprint (แบบมีน้ำหนัก เช่น เครื่องบิน / เรือ)**
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

## 📌 **4. Middleware ที่ใช้งาน**
| Middleware | รายละเอียด |
|------------|------------|
| `AssignRequestID` | เพิ่ม `X-Request-ID` ให้ทุก Request |
| `RequestLogger` | บันทึก Log ทุก API Request พร้อม Request ID |
| `CORS` | กำหนด `AllowOrigins` และ Headers ต่างๆ |
| `RateLimit` | จำกัดจำนวน Requests ต่อช่วงเวลา |
| `GZIPCompression` | เปิดใช้งาน GZIP เพื่อลดขนาด Response |
| `Recover` | จัดการ `panic` และคืนค่า Error Response |

---

## 📌 **5. Health Check API**
รองรับการตรวจสอบสถานะของระบบและ Database ตามมาตรฐาน Kubernetes
```sh
curl -X GET http://localhost:8080/live
curl -X GET http://localhost:8080/ready
```

---

## 📌 **6. ข้อผิดพลาดที่อาจพบและวิธีแก้ไข**
| ❌ Error | 🔹 สาเหตุ | ✅ วิธีแก้ไข |
|----------|----------|-------------|
| `relation \"emission_factors\" does not exist` | ไม่มี Table `emission_factors` ใน Database | ✅ รัน `migrate up` เพื่อสร้าง Table |
| `X-Request-ID ไม่แสดงใน Log` | Logger ไม่ถูกตั้งค่าให้ดึงค่า `request_id` | ✅ ตรวจสอบ Middleware `RequestLogger` และ `AssignRequestID` |
| `CORS Policy Error` | `AllowCredentials: true` แต่ `AllowOrigins: "*"` | ✅ ตั้งค่า `AllowCredentials: false` หรือกำหนด Origin ที่แน่นอน |
| `Swagger ไม่โหลด API Docs` | ไม่ได้ Import `_ "github.com/.../docs"` หรือไม่ได้เพิ่ม Route | ✅ ตรวจสอบ `main.go` ให้มี `app.Get("/swagger/*", swagger.HandlerDefault)` |
| `CI/CD Pipeline ล้มเหลว` | ไม่มีไฟล์ `go.mod` หรือใช้เวอร์ชัน Go ไม่ถูกต้อง | ✅ ตรวจสอบ `go mod tidy` และ Go เวอร์ชันใน Workflow |

---

📌 เปิด **Swagger UI** ที่ 👉 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---
fernserv@gmail.com
099-252-76O1