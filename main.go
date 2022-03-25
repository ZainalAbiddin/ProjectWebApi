package main

import (
	"net/http"

	"github.com/ZainalAbiddin/ProjectWebApi/controller"
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
	v1.GET("/mahasiswa", controller.GetDataMahasiswa)
	v1.POST("/mahasiswa", controller.CreateDataMahasiswa)
	v1.PUT("/mahasiswa/:nim", controller.UpdateDataMahasiswa)
	v1.DELETE("/mahasiswa", controller.DeleteDataMahasiswa)

	v1.GET("/matakuliah", controller.GetDataMatakuliah)
	v1.POST("/matakuliah", controller.CreateDataMatakuliah)
	v1.PUT("/matakuliah/:kode", controller.UpdateDataMaatakuliah)
	v1.DELETE("/matakuliah", controller.DeleteDataMatakuliah)
	r.Run()
}
