package models

import (
	"log"
	"time"

	"github.com/authentihub/api/database"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func Migrate() {
	db := database.GetDB()
	log.Println("🚀 Migrating models...")

	db.AutoMigrate(&User{})

	log.Println("✅ Migrated models")
}
