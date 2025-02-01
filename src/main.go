package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Takes in a JSON receipt and returns a JSON object with an ID
	router.POST("/receipts/process", func(c *gin.Context) {
		var receipt Receipt

		// bind the JSON payload with receipt
		if err := c.ShouldBindJSON(&receipt); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		log.Println(receipt)
	})

	// Start the webservice
	router.Run("127.0.0.1:8080")
}
