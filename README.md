# CarbonWize Digital Footprint Backend

🚀 **CarbonWize Digital Footprint Backend** เป็นระบบ API ที่ช่วยคำนวณ **Carbon Footprint** จากกิจกรรม เช่น **การขนส่ง (transportation)** โดยใช้ **Golang + Fiber + PostgreSQL** พร้อมรองรับ **Unit Test**

---

## 🔹 **1. Setup Project**
### ✅ **1.1 ติดตั้ง Dependency**
ให้รันคำสั่งนี้เพื่อโหลด package ทั้งหมดที่จำเป็น
```sh
go mod tidy
```

📌 **ปัญหาที่อาจเกิดขึ้น:**
- ❌ **error: go: no module directive in current directory** → ให้รัน `go mod init github.com/yourusername/carbonwize_digital_footprint_backend` ก่อน
- ❌ **error: unknown import path** → ตรวจสอบว่า `go.mod` มี dependencies ที่ถูกต้องและใช้ `go mod tidy` อัปเดต dependencies

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
📌 **ปัญหาที่อาจเกิดขึ้น:**
- ❌ **error: connection refused** → ตรวจสอบว่า PostgreSQL ทำงานอยู่ (`systemctl status postgresql` หรือ `pg_ctl status`)
- ❌ **error: password authentication failed** → ตรวจสอบ username/password ใน `config.yaml`

---

### ✅ **1.3 รัน Migration**
📌 **สร้างตารางและเพิ่มข้อมูลเริ่มต้น**
```sh
migrate -database "postgres://postgres:yourpassword@localhost:5432/carbon_db?sslmode=disable" -path migrations up
```
📌 **เช็คว่าตารางถูกสร้างแล้วหรือไม่**
```sh
psql -U postgres -d carbon_db -c "\dt"
```
📌 **ปัญหาที่อาจเกิดขึ้น:**
- ❌ **ERROR: relation "emission_factors" does not exist (SQLSTATE 42P01)** → Migration อาจไม่ถูกรัน ให้ใช้ `migrate up` อีกรอบ
- ❌ **no change** → ตารางอาจถูกสร้างไปแล้ว ลอง `migrate down` แล้ว `migrate up` ใหม่

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
📌 **ปัญหาที่อาจเกิดขึ้น:**
- ❌ **error: missing table `emission_factors`** → ตรวจสอบว่า migration ถูกรันแล้ว (`migrate up`)
- ❌ **error: port already in use** → พอร์ตอาจถูกใช้งานอยู่แล้ว ลองเปลี่ยนพอร์ตใน `config.yaml`

---

## 🔹 **2. ทดสอบ API ด้วย Postman**
📌 **คำนวณ Carbon Footprint (แบบไม่มีน้ำหนัก)**  
```json
POST http://localhost:8080/api/carbon/footprint/calculate/basic
{
  "activity_type": "transportation",
  "distance_km": 100,
  "vehicle_type": "car",
  "fuel_type": "gasoline"
}
```
📌 **คำนวณ Carbon Footprint (แบบใช้ `weight` เช่น เครื่องบิน / เรือ)**  
```json
POST http://localhost:8080/api/carbon/footprint/calculate/advanced
{
  "activity_type": "transportation",
  "distance_km": 1000,
  "vehicle_type": "airplane",
  "fuel_type": "jet_fuel",
  "weight": 80000
}
```
📌 **ปัญหาที่อาจเกิดขึ้น:**
- ❌ **error: record not found** → ฐานข้อมูลไม่มีค่าที่ตรงกับ request ให้ตรวจสอบข้อมูลด้วย `SELECT * FROM emission_factors;`
- ❌ **error: invalid input** → ตรวจสอบว่า request มีค่าที่ถูกต้อง เช่น `distance_km` ต้องเป็นตัวเลข

---

## 🔹 **3. รันทดสอบ (Unit Test)**
📌 **รันทุก Unit Test**
```sh
go test -v ./test
```
📌 **รันเฉพาะไฟล์ `service_test.go`**
```sh
go test -v ./test/service_test.go
```
📌 **รันทดสอบและเช็ค Coverage**
```sh
go test -cover ./test
```
📌 **ปัญหาที่อาจเกิดขึ้น:**
- ❌ **error: no test files** → ตรวจสอบว่าไฟล์ `_test.go` มีอยู่และมีฟังก์ชัน `TestXXX`
- ❌ **import cycle not allowed** → ตรวจสอบว่า test file ใช้ `package service_test` ไม่ใช่ `package service`

---

## 🎯 **สรุป**
| คำสั่ง | ใช้ทำอะไร |
|---------|-------------|
| `go mod tidy` | โหลด Dependency |
| `migrate up` | รัน Migration (สร้างตาราง) |
| `go run cmd/main.go` | รันเซิร์ฟเวอร์ |
| `air` | รันแบบ Hot Reload |
| `go test -v ./test` | รันทดสอบ Unit Test |

🔥 **ตอนนี้โปรเจกต์พร้อมใช้งานแล้ว! 🚀**

