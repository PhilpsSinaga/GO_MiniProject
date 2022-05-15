package model

import (
	"time"
)

type Book struct {
	Id                  int64  `json:"Id" form:"Id"`
	Nama_buku           string `json:"Nama_buku" form:"Nama_buku"`
	Penulis             string `json:"Penulis" form:"Penulis"`
	Penerbit            string `json:"Penerbit" form:"Penerbit"`
	Tahun_buku          int    `json:"Tahun_buku" form:"Tahun_buku"`
	Edisi               string `json:"Edisi" form:"Edisi"`
	Jumlah_halaman      int    `json:"Jumlah_halaman" form:"Jumlah_halaman"`
	Jumlah_halaman_baca int    `json:"Jumlah_halaman_baca" form:"Jumlah_halaman_baca"`
	Jenis_buku          string `json:"Jenis_buku" form:"Jenis_buku"`
	Jumlah_buku         int    `json:"Jumlah_buku" form:"Jumlah_buku"`
	Ketersedian_buku    string `json:"Ketersedian_buku" form:"Ketersediaan_buku"`
	Created_at          time.Time
	Updated_at          time.Time
}
