package main

import (
	"log"

	"github.com/authentihub/api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	models.Migrate()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the beginning of something great!",
		})
	})

	r.Run(":8000")
}
