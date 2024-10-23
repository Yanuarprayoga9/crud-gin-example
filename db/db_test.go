package db

import (
	"day1/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestConnectDb(t *testing.T) {
	dsn := "host=localhost user=postgres password=rahasia dbname=mydatabase sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Silent to avoid log pollution during tests
	})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get db instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		t.Fatalf("AutoMigrate failed: %v", err)
	}

	t.Log("Database connection successful and migrations ran")
}
