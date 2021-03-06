package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//koneksi database dari struct di kode golang
func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/kuliahdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("error koneksi ke dalam database")

	}
	db.AutoMigrate(&Mahasiswa{})
	db.AutoMigrate(&Matakuliah{})

	return db

}
