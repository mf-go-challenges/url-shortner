package main

import (
	"example.com/url-shortner/db"
	"example.com/url-shortner/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var (
	rdb     *redis.Client
	limiter *redis_rate.Limiter
)

func main() {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("cannot open log file: %v", err)
	}
	log.SetOutput(f)

	db.InitDB()

	server := gin.Default()
	server.Use(cors.Default())

	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	limiter = redis_rate.NewLimiter(rdb)

	handler.RegisterRoutes(server, limiter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
