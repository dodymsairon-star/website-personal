package main

import (
	"fintrack/config"
	"fintrack/routes"
	"log"
)

func main() {
	// Initialize database
	config.InitDB()
	defer config.CloseDB()

	// Setup router
	router := routes.SetupRouter()

	// Start server
	log.Println("🚀 Server starting on http://localhost:8080")
	log.Println("💾 Database: MySQL - Data tersimpan permanen!")
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}