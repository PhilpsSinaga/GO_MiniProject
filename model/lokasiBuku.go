package model

import "time"

type Location struct {
	Id         int64 `json:"Id" form:"Id"`
	Id_buku    int64 `json:"Id_buku" form:"Id_buku"`
	Lantai     int   `json:"Lantai" form:"Lantai"`
	Rak_buku   int   `json:"Rak_buku" form:"Rak_buku"`
	Baris_buku int   `json:"Baris_buku" form:"Baris_buku"`
	Created_at time.Time
	Updated_at time.Time
}
