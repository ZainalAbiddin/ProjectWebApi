package models

import "time"

type Matakuliah struct {
	Kode      string `json:"kode" gorm:"primary_key"`
	ID        int    `json:"id" `
	Nama      string `json:"nama"`
	Jumlah    int    `json:"jumlah"`
	Dosen     string `json:"dosen"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
**ID : int**
**Kode Matakuliah : string**
**Nama Mata Kuliah : string**
**Jumlah SKS : int**
**Dosen Pengampu : string**
 */
