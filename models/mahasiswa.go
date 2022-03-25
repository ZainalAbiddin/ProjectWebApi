package models

import "time"

type Mahasiswa struct {
	NIM            int    `json:"nim" gorm:"primary_key"`
	ID             int    `json:"id"`
	Nama           string `json:"nama"`
	Prodi          string `json:"prodi"`
	Fakultas       string `json:"fakultas"`
	Tahun_Angkatan int    `json:"tahun_angkatan"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

/*
1. **ID : int**
2. **Nama : string**
3. **Prodi : string**
4. **Fakultas : string**
5. **NIM : int**
6. **Tahun Angkatan : int**
*/
