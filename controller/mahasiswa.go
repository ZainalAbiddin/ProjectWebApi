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

type MahasiswaInput struct {
	NIM            int    `json:"nim" binding:"required,gt=9999"`
	ID             int    `json:"id"`
	Nama           string `json:"nama" binding:"required,gt=5"`
	Prodi          string `json:"prodi" binding:"required"`
	Fakultas       string `json:"fakultas" binding:"required"`
	Tahun_Angkatan int    `json:"tahun_angkatan" binding:"required"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// request isi data tabel
func GetDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
		"time": time.Now().Format(time.ANSIC),
	})

}

// create data
func CreateDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var mahasiswainput MahasiswaInput
	if err := c.ShouldBindJSON(&mahasiswainput); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.ActualTag() {
			case "required":
				errorMessage := fmt.Sprintf("Error %s, Pesan: Kolom %s harus di isi", e.Field(), e.Field())
				errorMessages = append(errorMessages, errorMessage)
			case "gt":
				errorMessage := fmt.Sprintf("Error %s, Pesan: Kolom %s tidak diisi sesuai format ", e.Field(), e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	//proses input data\

	mhs := models.Mahasiswa{
		NIM:            mahasiswainput.NIM,
		ID:             mahasiswainput.ID,
		Nama:           mahasiswainput.Nama,
		Prodi:          mahasiswainput.Prodi,
		Fakultas:       mahasiswainput.Fakultas,
		Tahun_Angkatan: mahasiswainput.Tahun_Angkatan,
	}

	// if err = mhs.Nama
	db.Create(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil input data",
		"data":    mhs,
		"time":    time.Now().Format(time.ANSIC),
	})
}

// update data
func UpdateDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data mahasiswa tak ada pak cik",
		})
		return
	}

	//validasi inputan
	var mahasiswainput MahasiswaInput
	if err := c.ShouldBindJSON(&mahasiswainput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// proses input update
	db.Model(&mhs).Update(mahasiswainput)

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil ubah data",
		"data":    mhs,
		"time":    time.Now().Format(time.ANSIC),
	})
}

// delete data
func DeleteDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Query("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data mahasiswa tak ada pak cik",
		})
		return
	}
	// proses hapus data
	db.Delete(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"Data":    true,
		"Message": "Berhasi hapus data",
	})
}
