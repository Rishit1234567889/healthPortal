#!/bin/bash

# This script runs database migrations

# Set default values
DB_USER=${PGUSER:-postgres}
DB_PASSWORD=${PGPASSWORD:-postgres}
DB_HOST=${PGHOST:-localhost}
DB_PORT=${PGPORT:-5432}
DB_NAME=${PGDATABASE:-hospital_portal}
MIGRATIONS_DIR="./migrations"

# Check if migrate tool is installed
if ! command -v migrate &> /dev/null; then
    echo "Error: migrate tool is not installed"
    echo "You can install it using: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
fi

# Build connection string
DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Perform migration based on command
case "$1" in
    up)
        echo "Running migrations up..."
        migrate -path ${MIGRATIONS_DIR} -database ${DB_URL} up
        ;;
    down)
        echo "Running migrations down..."
        migrate -path ${MIGRATIONS_DIR} -database ${DB_URL} down
        ;;
    create)
        if [ -z "$2" ]; then
            echo "Error: Missing migration name"
            echo "Usage: $0 create <migration_name>"
            exit 1
        fi
        echo "Creating new migration: $2"
        migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq "$2"
        ;;
    *)
        echo "Usage: $0 {up|down|create <name>}"
        exit 1
esac

echo "Migration command completed"
