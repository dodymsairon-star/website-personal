package routes

import (
	"fintrack/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Serve static files
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// API routes
	api := r.Group("/api")
	{
		api.GET("/dashboard", handlers.GetDashboardData)
		api.GET("/stats", handlers.GetStats)
		api.GET("/category-stats", handlers.GetCategoryStats)
		api.GET("/monthly", handlers.GetMonthlyData)

		api.GET("/transactions", handlers.GetAllTransactions)
		api.GET("/transactions/search", handlers.SearchTransactions)
		api.GET("/transactions/filter", handlers.FilterTransactions)
		api.GET("/transactions/:id", handlers.GetTransaction)
		api.POST("/transactions", handlers.CreateTransaction)
		api.PUT("/transactions/:id", handlers.UpdateTransaction)
		api.DELETE("/transactions/:id", handlers.DeleteTransaction)

		api.GET("/export/csv", handlers.ExportCSV)
	}

	return r
}