package db

import (
	"fmt"
	"log"

	"github.com/Naomejoy/app-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database successfully")

	// Auto Migrate
	// log.Println("Running AutoMigrate...")
	// err = DB.AutoMigrate(
	// 	&domain.Application{},
	// 	&domain.ApplicationStatus{},
	// 	&domain.ApplicationUploadedFileType{},
	// )
	// if err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }
	log.Println("Database migration completed successfully")
}
