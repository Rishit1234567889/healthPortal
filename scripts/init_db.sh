#!/bin/bash

# This script initializes the database with the initial schema


# Set environment variables from DATABASE_URL
if [ -n "$DATABASE_URL" ]; then
    # Extract parts from DATABASE_URL
    # Format: postgres://username:password@hostname:port/database
    DB_URL=${DATABASE_URL#postgres://}
    DB_USER=$(echo $DB_URL | cut -d':' -f1)
    DB_PASSWORD=$(echo $DB_URL | cut -d':' -f2 | cut -d'@' -f1)
    DB_HOST=$(echo $DB_URL | cut -d'@' -f2 | cut -d':' -f1)
    DB_PORT=$(echo $DB_URL | cut -d':' -f3 | cut -d'/' -f1)
    DB_NAME=$(echo $DB_URL | cut -d'/' -f2)
else
    # Use environment variables directly
    DB_USER=${PGUSER}
    DB_PASSWORD=${PGPASSWORD}
    DB_HOST=${PGHOST}
    DB_PORT=${PGPORT}
    DB_NAME=${PGDATABASE}
fi

# Check if psql is installed
if ! command -v psql &> /dev/null; then
    echo "Error: psql is not installed"
    exit 1
fi

echo "Initializing database..."
echo "Host: $DB_HOST"
echo "Port: $DB_PORT"
echo "Database: $DB_NAME"
echo "User: $DB_USER"

# Create the schema
PGPASSWORD=${DB_PASSWORD} psql -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} -d ${DB_NAME} -c "
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
CREATE INDEX IF NOT EXISTS idx_patients_name ON patients(name);

-- Create index on patients phone_number for faster search
CREATE INDEX IF NOT EXISTS idx_patients_phone ON patients(phone_number);

-- Insert initial admin users if they don't exist
INSERT INTO users (name, email, password, role)
SELECT 'Admin Doctor', 'doctor@example.com', '\$2a\$10\$NqRvFBJbhYY5XKHMfA9XJu2dTd5QPjJhPfUo5zuYmOW.mSZ9ThAGu', 'doctor'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'doctor@example.com');

INSERT INTO users (name, email, password, role)
SELECT 'Admin Receptionist', 'receptionist@example.com', '\$2a\$10\$NqRvFBJbhYY5XKHMfA9XJu2dTd5QPjJhPfUo5zuYmOW.mSZ9ThAGu', 'receptionist'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'receptionist@example.com');
"

# Seed sample data
echo "Seeding database with sample data..."
PGPASSWORD=${DB_PASSWORD} psql -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} -d ${DB_NAME} -c "
-- Insert sample patients if they don't exist
INSERT INTO patients (name, age, gender, address, phone_number, medical_history, diagnosis, treatment, notes)
SELECT 'John Doe', 45, 'male', '123 Main St, Anytown', '555-123-4567', 'Hypertension, Diabetes', 'Flu', 'Rest and fluids', 'Patient should follow up in 1 week'
WHERE NOT EXISTS (SELECT 1 FROM patients WHERE name = 'John Doe' AND phone_number = '555-123-4567');

INSERT INTO patients (name, age, gender, address, phone_number, medical_history, diagnosis, treatment, notes)
SELECT 'Jane Smith', 32, 'female', '456 Oak Ave, Somewhere', '555-987-6543', 'Asthma', 'Bronchitis', 'Antibiotics, inhaler', 'Improving slowly'
WHERE NOT EXISTS (SELECT 1 FROM patients WHERE name = 'Jane Smith' AND phone_number = '555-987-6543');

INSERT INTO patients (name, age, gender, address, phone_number, medical_history, diagnosis, treatment, notes)
SELECT 'Michael Johnson', 28, 'male', '789 Pine Rd, Nowhere', '555-456-7890', 'None', 'Sprained ankle', 'RICE therapy', 'Use crutches for 2 weeks'
WHERE NOT EXISTS (SELECT 1 FROM patients WHERE name = 'Michael Johnson' AND phone_number = '555-456-7890');

INSERT INTO patients (name, age, gender, address, phone_number, medical_history, diagnosis, treatment, notes)
SELECT 'Emily Davis', 65, 'female', '321 Elm Blvd, Anyplace', '555-789-0123', 'Arthritis, High cholesterol', 'Pneumonia', 'Antibiotics, rest', 'Monitor oxygen levels'
WHERE NOT EXISTS (SELECT 1 FROM patients WHERE name = 'Emily Davis' AND phone_number = '555-789-0123');
"

echo "Database initialization completed"