package db

import (
	"day1/model/domain"
	"log"
	"os"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	maxRetries := 5
	retryDelay := 5 * time.Second

	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database. Retrying in %v... (%d/%d)\n", retryDelay, i+1, maxRetries)
		time.Sleep(retryDelay)
	}

	if err != nil {
		log.Fatal("Failed to connect to database after maximum retries. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&domain.User{})

	DB = Dbinstance{
		Db: db,
	}
	log.Println("connected to database")
	return db
}
