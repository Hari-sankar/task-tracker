# task-tracker
# Task Tracker API

A robust REST API for task management built with Go, featuring user authentication, task operations, and comprehensive error handling.

## Features
- User Authentication with JWT
- Task Management (CRUD operations)
- PostgreSQL Database
- Swagger Documentation
- Docker Support
- Structured Error Handling
- Request Validation
- Database Migrations
- Zap Logger Integration

## Tech Stack
- Go 1.21+
- Gin Web Framework
- PostgreSQL
- Docker & Docker Compose
- Swagger
- JWT Authentication
- Zap Logger

## Setup Instructions

### Using Docker
1. Clone the repository
```bash
git clone https://github.com/Hari-sankar/task-tracker
cd task-tracker
```
2. Copy .env.example to .env
```bash
cp .env.example .env
```
3. Build and run the Docker containers
```bash
docker-compose up --build
```
### Without Docker
1. Clone the repository
```bash
git clone https://github.com/Hari-sankar/task-tracker
cd task-tracker
```
2. Install dependencies
```bash
go mod tidy
```
3. Set up the database
```bash
make migrate-up
```
4. Run the application
```bash
go run main.go start
```
# Environment Variables Configuration

## Database Configuration
```env
# PostgreSQL User
DB_USER=postgres

# PostgreSQL Password
DB_PASSWORD=password123

# Database Name
DB_NAME=task_tracker

# Database Connection String
# For Docker:
DB_SOURCE=postgresql://postgres:password123@postgres:5432/task_tracker?sslmode=disable

# For Local Development:
# DB_SOURCE=postgresql://postgres:password123@localhost:5432/task_tracker?sslmode=disable
```

# API Documentation

## Swagger Documentation
The API documentation is available in multiple formats:

### Interactive Swagger UI
Access the interactive API documentation at:
```http
http://localhost:3000/swagger/index.html
```

## Raw Documentation Files

The API specification is available in two formats in the `/docs` folder:

- **JSON Format**: [`/docs/swagger.json`](/docs/swagger.json)
- **YAML Format**: [`/docs/swagger.yaml`](/docs/swagger.yaml)

These files can be imported into any Swagger/OpenAPI compatible tool for viewing or testing the API endpoints.


### View Documentation Without Starting Server
1. Open [Swagger Editor](https://editor.swagger.io/)
2. Import either of these files from the `/docs` folder:
   - [`swagger.json`](/docs/swagger.json)
   - [`swagger.yaml`](/docs/swagger.yaml)