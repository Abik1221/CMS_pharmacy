package user

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleSuperAdmin     Role = "SUPER_ADMIN"
	RoleSubAdmin       Role = "SUB_ADMIN"
	RolePharmacyOwner  Role = "PHARMACY_OWNER"
	RolePharmacyStaff  Role = "PHARMACY_STAFF"
	RoleCashier        Role = "CASHIER"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100)" json:"name"`
	Email     string         `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
	Password  string         `gorm:"type:text" json:"-"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Role      Role           `gorm:"type:varchar(30)" json:"role"`
	PharmacyID *uint         `json:"pharmacy_id,omitempty"` // nil for platform admins
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

