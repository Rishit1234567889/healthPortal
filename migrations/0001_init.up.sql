-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('doctor', 'receptionist')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create patients table

CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL,
    gender VARCHAR(50) NOT NULL,
    address TEXT NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    medical_history TEXT,
    diagnosis TEXT,
    treatment TEXT,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create index on patients name for faster search
CREATE INDEX idx_patients_name ON patients(name);

-- Create index on patients phone_number for faster search
CREATE INDEX idx_patients_phone ON patients(phone_number);

-- Insert initial admin users

INSERT INTO users (name, email, password, role)
VALUES 
    ('Admin Doctor', 'doctor@example.com', '$2a$10$NqRvFBJbhYY5XKHMfA9XJu2dTd5QPjJhPfUo5zuYmOW.mSZ9ThAGu', 'doctor'),  -- password: password123
    ('Admin Receptionist', 'receptionist@example.com', '$2a$10$NqRvFBJbhYY5XKHMfA9XJu2dTd5QPjJhPfUo5zuYmOW.mSZ9ThAGu', 'receptionist'); -- password: password123
