package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func Connect() *gorm.DB {
	url := os.Getenv("DATABASE_URL")

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Database not ready yet (attempt %d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if db == nil {
		log.Fatal("Database failed to connect")
		return nil
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Error connecting to database", err.Error())
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected to database")

	return db
}
