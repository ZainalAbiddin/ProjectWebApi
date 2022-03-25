package main

import (
	"net/http"

	"github.com/ZainalAbiddin/ProjectWebApi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// models
	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Pesan": "Golang Web Api Sederhana",
		})
	})
	r.Run()
}
