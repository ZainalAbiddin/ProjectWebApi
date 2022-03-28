package models

import "time"

type Mahasiswa struct {
	NIM            int    `json:"nim"`
	ID             int    `json:"id" gorm:"primary_key AUTO_INCREMENT NOT_NULL"`
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
