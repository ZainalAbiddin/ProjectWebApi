package controller

import (
	"net/http"
	"time"

	"github.com/ZainalAbiddin/ProjectWebApi/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MatakuliahInput struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Kode      string `json:"kode"`
	Nama      string `json:"nama"`
	Jumlah    int    `json:"jumlah"`
	Dosen     string `json:"dosen"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetDataMatakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mtk []models.Matakuliah
	db.Find(&mtk)
	c.JSON(http.StatusOK, gin.H{
		"data": mtk,
		"time": time.Now(),
	})

}
