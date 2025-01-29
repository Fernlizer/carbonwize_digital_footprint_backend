CREATE TABLE IF NOT EXISTS emission_factors (
    id SERIAL PRIMARY KEY,
    activity_type VARCHAR(50) NOT NULL,  -- เช่น 'transportation' หรือ 'energy_consumption'
    vehicle_type VARCHAR(50) NOT NULL,   -- เช่น 'car', 'bus', 'motorcycle' (ถ้าเป็น transportation)
    fuel_type VARCHAR(50) NOT NULL,      -- เช่น 'gasoline', 'diesel', 'electric'
    emission_factor NUMERIC NOT NULL,    -- ค่าการปล่อย CO2 (kg/km หรือ kg/kWh)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


-- กำหนด Default Data ข้อมูลอ้างอิงจาก
-- EPA (Environmental Protection Agency) และ IPCC (Intergovernmental Panel on Climate Change)
INSERT INTO emission_factors (activity_type, vehicle_type, fuel_type, emission_factor) VALUES
    -- รถยนต์ทั่วไป
    ('transportation', 'car', 'gasoline', 0.254),  -- kg CO2/km
    ('transportation', 'car', 'diesel', 0.270),    -- kg CO2/km
    ('transportation', 'electric_car', 'electric', 0.050), -- kg CO2/km (รวมจากไฟฟ้า)

    -- รถมอเตอร์ไซค์
    ('transportation', 'motorcycle', 'gasoline', 0.120), -- kg CO2/km
    ('transportation', 'motorcycle', 'electric', 0.030), -- kg CO2/km

    -- รถบัส
    ('transportation', 'bus', 'diesel', 0.080),   -- kg CO2/km (ต่อผู้โดยสาร)
    ('transportation', 'bus', 'electric', 0.020), -- kg CO2/km (ต่อผู้โดยสาร)

    -- รถไฟ
    ('transportation', 'train', 'diesel', 0.040),    -- kg CO2/km (ต่อผู้โดยสาร)
    ('transportation', 'train', 'electric', 0.010),  -- kg CO2/km (ต่อผู้โดยสาร)

    -- เรือโดยสาร
    ('transportation', 'ship', 'diesel', 0.300),   -- kg CO2/km (ต่อผู้โดยสาร)
    ('transportation', 'ship', 'electric', 0.100), -- kg CO2/km (ต่อผู้โดยสาร)

    -- เครื่องบิน
    ('transportation', 'airplane', 'jet_fuel', 0.250); -- kg CO2/km (ต่อผู้โดยสาร)
