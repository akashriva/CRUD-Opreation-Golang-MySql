package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading environment variables: ", err)
	}

	// Retrieve the PORT from the environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new Gin router
	router := gin.Default()
	fmt.Println("********")
	fmt.Println("********SERVER*****")
	fmt.Println("**************IS*****")
	fmt.Println("****************Connected*****")
	fmt.Println("*************************ON*****")
	fmt.Printf("****************************PORT %v ***\n", port)
	// Define your routes and handlers here

	// Run the server on the specified PORT
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

}

func loadEnv() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
