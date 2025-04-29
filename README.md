# CMS_pharmacy
This is a content management system backend infrastructure build using golang for Pharmacy in Africa.

## The entire proper folder structure is the fallowing.

/cmd/
├── main.go                 # Starts the Gin server
├── migrate/
│   └── main.go             # (Optional) Go-based runner for migrations

/config/
├── config.go               # App configs (env, port, keys)
├── db.go                   # DB initialization (GORM)
├── env.go                  # Load .env variables

/internal/
├── auth/                   # Authentication logic
│   ├── handler.go          # Login, register controllers
│   ├── service.go          # Auth business logic
│   ├── jwt.go              # JWT generation and parsing
│   └── middleware.go       # Auth and role-based access middleware

├── user/                   # User and role management
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go

├── pharmacy/               # Pharmacy-specific features
│   ├── handler.go          # Inventory, medicine, reports
│   ├── service.go
│   ├── repository.go
│   └── model.go

├── cashier/                # Cashier-specific functionality
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go

├── subscription/           # Subscription plans, payments
│   ├── handler.go
│   ├── service.go
│   ├── payment.go          # Stripe or Chapa integration
│   ├── repository.go
│   └── model.go

├── chatbot/                # MedChatbot logic (Premium only)
│   ├── handler.go
│   ├── service.go
│   └── integration.go

├── patient/                # Patient and prescription management
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go

├── analytics/              # Advanced reporting and dashboards
│   ├── handler.go
│   ├── service.go
│   └── model.go

├── routes/                 # Central route registration
│   └── routes.go

/migrations/                # SQL-based migrations
├── 001_create_users.up.sql
├── 001_create_users.down.sql
├── 002_create_pharmacies.up.sql
├── ...
└── README.md               # Docs for how to run migrations

/pkg/
├── jwt/                    # JWT utilities for token signing/parsing
│   └── jwt.go
├── logger/                 # Custom logger (e.g., Zap or Logrus)
│   └── logger.go
├── utils/                  # Utility functions (string helpers, etc.)
│   └── helpers.go

/public/                    # Static files (e.g., Swagger docs, uploads)

/docs/                      # API and system documentation (Swagger)
/tests/                     # Unit and integration tests
/go.mod
/go.sum
/.env                       # Environment variables
