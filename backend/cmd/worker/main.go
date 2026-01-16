package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"social-media-scheduler/internal/db"
	"social-media-scheduler/internal/queue"
	"social-media-scheduler/internal/worker"
	"syscall"
	"time"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initiate Database
	database := db.Connect()

	// Initiate redis connection
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb := os.Getenv("REDIS_DB")
	redisClient := queue.NewRedisClient(redisAddress, redisPassword, redisDb)
	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	log.Printf("Redis Connection Established Successfully")

	w := worker.NewWorker(redisClient, database)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping worker...")
			return
		case <-ticker.C:
			err := w.Run(ctx)
			if err != nil {
				log.Println("Error worker: ", err)
			}
		}
	}
}
