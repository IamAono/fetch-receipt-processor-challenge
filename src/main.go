package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	id := 1
	// a map where the key is the id and the value is the amount of points earned
	points := make(map[int]int)
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

		// calculate the amount of points that were earned from the recipt and save the value in a map
		points[id] = receipt.calcPoints()
		log.Println("Total points:", points[id])

		c.JSON(http.StatusOK, gin.H{"id": id})
		id++
	})
	router.GET("/receipts/:id/points", func(c *gin.Context) {
		idReq, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		// check if the id being requested exists in the map
		_, ok := points[idReq]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		}
		c.JSON(http.StatusOK, gin.H{"points": points[idReq]})

	})

	// Start the webservice
	router.Run("127.0.0.1:8080")
}
