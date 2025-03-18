# GoTrack API – Geolocation-powered travel backend 🚀

## Overview
Travel Backend is a high-performance backend service built with **Golang** and **PostgreSQL**, designed to power an upcoming travel application. It leverages the **Fiber framework** for efficient request handling, **Gorm** for ORM, and integrates JWT authentication, Swagger documentation, and AWS SDK for enhanced capabilities.

## Tech Stack ⚙️
- **[Fiber](https://gofiber.io/)** - Fast and lightweight web framework for Golang
- **[JWT](https://github.com/gofiber/jwt)** - Middleware for authentication and authorization
- **[Swagger](https://github.com/gofiber/swagger)** - API documentation tool
- **[AWS SDK for Go](https://github.com/aws/aws-sdk-go)** - AWS service integrations
- **[Gorm](https://gorm.io/)** - ORM for Golang
- **[lib/pq](https://github.com/lib/pq)** - PostgreSQL driver
- **[Godotenv](https://github.com/joho/godotenv)** - Environment variable management
- **[Gocron](https://github.com/go-co-op/gocron)** - Job scheduling library
- **[Zap](https://github.com/uber-go/zap)** - High-performance logging
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Password hashing for security
- **[Air](https://github.com/cosmtrek/air)** - Live reloading for development

## Features 🔥
- **Fast API responses** with Fiber
- **Secure authentication** with JWT
- **API documentation** with Swagger
- **PostgreSQL integration** with Gorm
- **Environment configuration** with Godotenv
- **Automated task scheduling** with Gocron
- **Structured logging** with Zap
- **Live development reloading** with Air
- **Secure password handling** with Bcrypt

## Setup & Installation 🛠️
1. **Clone the repository:**
   ```sh
   git clone <repository-url>
   cd travel_backend
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Set up environment variables:**
   Create a `.env` file and configure necessary variables like database URL, AWS credentials, etc.
    ```sh
    ENV=Development  # or "Production"
    # PostgreSQL Database Configuration
    DB_USER=
    DB_PASSWORD=
    DB_NAME=
    DB_HOST=
    DB_PORT=5432
    # AWS Credentials
    AWS_ACCESS_KEY_ID=
    AWS_SECRET_ACCESS_KEY=
    AWS_REGION=
    # Additional
    # Example: REDIS_URL=

4. **Run the application:**
   ```sh
   go run main.go
   ```

## API Documentation 📄
Once the server is running, visit `http://localhost:3000/swagger/index.html` to explore the API docs.

## License 📝
This project is licensed under the **MIT License**.
