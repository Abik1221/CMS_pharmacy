package config

import (
	"log"

	"github.com/abik1221/pharmacy/internal/user"
	"github.com/abik1221/pharmacy/internal/pharmacy"
	"github.com/abik1221/pharmacy/internal/cashier"
	"github.com/abik1221/pharmacy/internal/subscription"
	"github.com/abik1221/pharmacy/internal/patient"
	"github.com/abik1221/pharmacy/internal/analytics"
)

// AutoMigrate registers all models with GORM for migration
func AutoMigrate() {
	err := DB.AutoMigrate(
		&user.User{},
		&user.Role{},
		&pharmacy.Medicine{},
		&pharmacy.Inventory{},
		&cashier.Sale{},
		&subscription.Subscription{},
		&subscription.Plan{},
		&patient.Patient{},
		&analytics.Report{},
	)
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	log.Println("📦 Auto migration completed successfully.")
}
