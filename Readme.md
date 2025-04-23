# 🏥 Hospital Management System API

![Go](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-✓-2496ED?logo=docker)

A secure REST API for hospital management with role-based access control (Receptionists & Doctors).

## ✨ Features

- **🔐 Authentication**
  - JWT-based authentication
  - Role-based access (Receptionist/Doctor)
  - Password hashing with bcrypt

- **👥 Patient Management**
  - Patient registration with auto-generated IDs
  - Medical history tracking
  - Search and filtering capabilities

- **📊 Medical Records**
  - Diagnosis and treatment records
  - Prescription management
  - Doctor-patient association

- **⚙️ System Features**
  - Auto database migrations
  - Structured logging (Zap)
  - Configurable CORS policies
  - Swagger API documentation
  - Dockerized deployment

## 🛠 Tech Stack

| Category       | Technology                          |
|----------------|-------------------------------------|
| **Backend**    | Go 1.19+, Gin, GORM                 |
| **Database**   | PostgreSQL                          |
| **Auth**       | JWT, bcrypt                         |
| **Logging**    | Zap                                 |
| **Config**     | Viper                               |
| **DevOps**     | Docker, Docker Compose              |
| **Testing**    | testify, httptest                   |

## 📂 Project Structure

```text
hospital-management-system/
├── cmd/               # Main application
├── configs/           # Configuration files
├── internal/          # Core application logic
│   ├── auth/          # Authentication
│   ├── controllers/   # Request handlers  
│   ├── middlewares/   # Gin middleware
│   ├── models/        # Database schema
│   ├── repositories/  # Data access
│   ├── routes/        # API endpoints
│   ├── services/      # Business logic
│   └── utils/         # Helpers
├── migrations/        # DB migrations
├── scripts/          # Deployment scripts
└── tests/            # Integration tests
```


## Setup and Installation


## 🚀 Installation

### Prerequisites
- Docker & Docker Compose
- Go 1.19+ (for development)

### With Docker (Recommended)


git clone https://github.com/RishitRishit1234567889/hospital-management-system.git
cd hospital-management-system

# Start services
docker-compose up -d

# Setup database
docker-compose run --rm migrate
docker-compose run --rm seed

### With Docker (Recommended) 

# Install dependencies
go mod download

# Set up PostgreSQL and update configs/config.yaml
# Run migrations
go run cmd/server/main.go --migrate

# Start server
go run cmd/server/main.go

## Usage

1. Open a web browser and navigate to `http://localhost:8080`
2. Log in as either a doctor or receptionist
3. Use the respective portal to manage patient records

## API Documentation

A Postman collection is provided in the repository for testing the API endpoints.