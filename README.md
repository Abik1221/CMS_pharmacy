# CMS_pharmacy
This is a content management system backend infrastructure build using golang for Pharmacy in Africa.

This system is currently on its way to be patanted! Copying or ownng will gonna make you Punished.


## The entire proper folder structure is the fallowing

# Pharmacy Management System - Backend

A structured Go backend project using Gin and GORM, following Clean Architecture principles.

## Project Structure

### `/cmd/`
- **Main Application**
  - `main.go` - Starts the Gin http server
- **Migrations Runner**
  - `/migrate/main.go` - Optional Go-based runner for DB migrations

### `/config/`
- Configuration Management
  - `config.go` - App configuration (env, port, keys)
  - `db.go` - GORM DB initialization
  - `env.go` - Loads .env variables

### `/internal/`
#### Core Application Components

- **Authentication** (`/auth/`)
  - `handler.go` - Login/register handlers
  - `service.go` - Auth business logic
  - `jwt.go` - JWT generation & parsing
  - `middleware.go` - Auth + RBAC middleware

- **User Management** (`/user/`)
  - `handler.go` - User endpoints
  - `service.go` - User business logic
  - `repository.go` - Database operations
  - `model.go` - Data models

- **Pharmacy Operations** (`/pharmacy/`)
  - `handler.go` - Inventory, medicines, reports
  - `service.go` - Business logic
  - `repository.go` - Data access
  - `model.go` - Domain models

- **Cashier System** (`/cashier/`)
  - `handler.go` - Transaction endpoints
  - `service.go` - Sales logic
  - `repository.go` - Data operations
  - `model.go` - Data structures

- **Subscriptions** (`/subscription/`)
  - `handler.go` - Plan management
  - `service.go` - Subscription logic
  - `payment.go` - Stripe/Chapa integration
  - `repository.go` - Data layer
  - `model.go` - Subscription models

- **Chatbot** (`/chatbot/`)
  - `handler.go` - Chat endpoints
  - `service.go` - Bot logic
  - `integration.go` - External service integration

- **Patient Management** (`/patient/`)
  - `handler.go` - Patient endpoints
  - `service.go` - Prescription logic
  - `repository.go` - Data access
  - `model.go` - Patient models

- **Analytics** (`/analytics/`)
  - `handler.go` - Reports endpoints
  - `service.go` - Analytics logic
  - `model.go` - Data models

- **Routing** (`/routes/`)
  - `routes.go` - Registers all API endpoints

### `/migrations/`
- Database Schema Management
  - SQL migration files (e.g., `001_create_users.up.sql`)
  - `README.md` - Migration instructions

### `/pkg/`
- Shared Utilities
  - **JWT** (`/jwt/jwt.go`) - Token utilities
  - **Logging** (`/logger/logger.go`) - Custom logger
  - **Utilities** (`/utils/helpers.go`) - Common helpers

### Supporting Directories
- `/public/` - Static files (Swagger docs, uploads)
- `/docs/` - API documentation & specs
- `/tests/` - Unit & integration tests

## Root Files
- `.env` - Environment variables
- `go.mod` - Go dependencies
- `go.sum` - Dependency checksums

