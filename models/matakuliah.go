package models

import "time"

type Matakuliah struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Kode      string `json:"kode"`
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
