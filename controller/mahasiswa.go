package controller

import (
	"net/http"
	"time"

	"github.com/ZainalAbiddin/ProjectWebApi/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	ID             int    `json:"id" gorm:"primary_key"`
	Nama           string `json:"nama"`
	Prodi          string `json:"prodi"`
	Fakultas       string `json:"fakultas"`
	NIM            int    `json:"nim"`
	Tahun_Angkatan int    `json:"tahun_angkatan"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func GetData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
		"time": time.Now(),
	})

}
