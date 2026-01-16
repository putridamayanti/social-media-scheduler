package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"social-media-scheduler/internal/db"
	"social-media-scheduler/internal/handlers"
	"social-media-scheduler/internal/middleware"
	"social-media-scheduler/internal/queue"
	"social-media-scheduler/internal/repositories"
	"social-media-scheduler/internal/services"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	log.Println("Database", os.Getenv("DATABASE_URL"))

	database := db.Connect()

	// Auto Migrate
	err := db.Migrate(database)
	if err != nil {
		log.Fatal("Migration failed", err)
	}

	// Initiate redis connection
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb := os.Getenv("REDIS_DB")
	redisClient := queue.NewRedisClient(redisAddress, redisPassword, redisDb)
	err = redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	// Queue
	qu := queue.NewScheduler(redisClient)

	// Define handlers
	authRepository := repositories.NewAuthRepository(database)
	authService := services.NewAuthService(authRepository, repositories.NewUserRepository(database))
	authHandler := handlers.NewAuthHandler(authService)

	postService := services.NewPostService(repositories.NewPostRepository(database))
	postHandler := handlers.NewPostHandler(postService, qu)

	userService := services.NewUserService(repositories.NewUserRepository(database))
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)

		protected := api.Group("/")
		{
			protected.Use(middleware.AuthMiddleware(authRepository))

			protected.POST("/posts", postHandler.Create)
			protected.GET("/posts", postHandler.GetAll)
			protected.GET("/posts/:id", postHandler.GetByID)
			protected.PUT("/posts/:id", postHandler.Update)
			protected.DELETE("/posts/:id", postHandler.Delete)

			protected.POST("/users", userHandler.Create)
			protected.GET("/users", userHandler.GetAll)
			protected.GET("/users/:id", userHandler.GetByID)
			protected.PUT("/users/:id", userHandler.Update)
			protected.DELETE("/users/:id", userHandler.Delete)

			protected.DELETE("/logout", authHandler.Logout)
		}

	}

	port := os.Getenv("PORT")
	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
