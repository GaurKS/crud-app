package main

import (
	"log"
	"os"

	db "github.com/GaurKS/crud-app/pkg/config"
	"github.com/GaurKS/crud-app/pkg/routes"
	"github.com/GaurKS/crud-app/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Setting variables for app env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection and handler assignment
	DB := db.Init()
	h := services.New(DB)
		
	// Initialize Gin-Gonic router
	router := gin.Default()
	router.Use(gin.Logger())
	r := router.Group("/api")
	routes.TodoRouter(r.Group("/todo"), &h)
	router.Run(":" + os.Getenv("PORT"))
}