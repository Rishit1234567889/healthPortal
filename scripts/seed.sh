#!/bin/bash

# This script seeds the database with initial data

# Set default values
DB_USER=${PGUSER:-postgres}
DB_PASSWORD=${PGPASSWORD:-postgres}
DB_HOST=${PGHOST:-localhost}
DB_PORT=${PGPORT:-5432}
DB_NAME=${PGDATABASE:-hospital_portal}

# Check if psql is installed
if ! command -v psql &> /dev/null; then
    echo "Error: psql is not installed"
    exit 1
fi

# Sample patients data
read -r -d '' SEED_SQL << EOM
-- Insert sample patients
INSERT INTO patients (name, age, gender, address, phone_number, medical_history, diagnosis, treatment, notes)
VALUES
    ('John Doe', 45, 'male', '123 Main St, Anytown', '555-123-4567', 'Hypertension, Diabetes', 'Flu', 'Rest and fluids', 'Patient should follow up in 1 week'),
    ('Jane Smith', 32, 'female', '456 Oak Ave, Somewhere', '555-987-6543', 'Asthma', 'Bronchitis', 'Antibiotics, inhaler', 'Improving slowly'),
    ('Michael Johnson', 28, 'male', '789 Pine Rd, Nowhere', '555-456-7890', 'None', 'Sprained ankle', 'RICE therapy', 'Use crutches for 2 weeks'),
    ('Emily Davis', 65, 'female', '321 Elm Blvd, Anyplace', '555-789-0123', 'Arthritis, High cholesterol', 'Pneumonia', 'Antibiotics, rest', 'Monitor oxygen levels');
EOM

# Run the seed SQL
echo "Seeding database with sample data..."
PGPASSWORD=${DB_PASSWORD} psql -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} -d ${DB_NAME} -c "${SEED_SQL}"

echo "Database seeded successfully"
