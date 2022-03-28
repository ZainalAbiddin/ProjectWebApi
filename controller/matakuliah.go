package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ZainalAbiddin/ProjectWebApi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type MatakuliahInput struct {
	Kode      string `json:"kode" binding:"required"`
	ID        int    `json:"id"`
	Nama      string `json:"nama" binding:"required,gt=3"`
	Jumlah    int    `json:"jumlah" binding:"required"`
	Dosen     string `json:"dosen" binding:"required"`
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
		"time": time.Now().Format(time.ANSIC),
	})

}

// create data
func CreateDataMatakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var matakuliahinput MatakuliahInput
	if err := c.ShouldBindJSON(&matakuliahinput); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.ActualTag() {
			case "required":
				errorMessage := fmt.Sprintf("Error %s, Pesan: Kolom %s harus di isi", e.Field(), e.Field())
				errorMessages = append(errorMessages, errorMessage)
			case "gt":
				errorMessage := fmt.Sprintf("Error %s, Pesan: Kolom %s tidak diisi sesuai format", e.Field(), e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
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
		"time":    time.Now().Format(time.ANSIC),
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
	var matakuliahinput MatakuliahInput
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
		"time":    time.Now().Format(time.ANSIC),
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
