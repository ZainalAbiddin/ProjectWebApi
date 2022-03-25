package controller

import (
	"net/http"
	"time"

	"github.com/ZainalAbiddin/ProjectWebApi/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MatakuliahInput struct {
	Kode      string `json:"kode" gorm:"primary_key"`
	ID        int    `json:"id" `
	Nama      string `json:"nama"`
	Jumlah    int    `json:"jumlah"`
	Dosen     string `json:"dosen"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// request isi data tabel
func GetDataMatakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mtk []models.Matakuliah
	db.Find(&mtk)
	c.JSON(http.StatusOK, gin.H{
		"data": mtk,
		"time": time.Now(),
	})

}

// create data
func CreateDataMatakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var matakuliahinput MatakuliahInput
	if err := c.ShouldBindJSON(&matakuliahinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//proses input data
	mtk := models.Matakuliah{
		Kode:   matakuliahinput.Kode,
		ID:     matakuliahinput.ID,
		Nama:   matakuliahinput.Nama,
		Jumlah: matakuliahinput.Jumlah,
		Dosen:  matakuliahinput.Dosen,
	}

	db.Create(&mtk)

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil input data",
		"data":    mtk,
		"time":    time.Now(),
	})
}

// update data
func UpdateDataMaatakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek data
	var mtk models.Matakuliah
	if err := db.Where("kode = ?", c.Param("kode")).First(&mtk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data matakuliah tak ada pak cik",
		})
		return
	}

	//validasi inputan
	var matakuliahinput MahasiswaInput
	if err := c.ShouldBindJSON(&matakuliahinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// proses input update
	db.Model(&mtk).Update(matakuliahinput)

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil ubah data",
		"data":    mtk,
		"time":    time.Now(),
	})
}

// delete data
func DeleteDataMatakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek data
	var mtk models.Matakuliah
	if err := db.Where("kode = ?", c.Query("kode")).First(&mtk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data matakuliah tak ada pak cik",
		})
		return
	}
	// proses hapus data
	db.Delete(&mtk)

	c.JSON(http.StatusOK, gin.H{
		"Data":    true,
		"Message": "Berhasi hapus data",
	})
}
