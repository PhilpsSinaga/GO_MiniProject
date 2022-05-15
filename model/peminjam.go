package model

import "time"

type Borower struct {
	Id                int64 `json:"Id" form:"Id"`
	Id_buku           int64 `json:"Id_buku" form:"Id_buku"`
	Tgl_peminjaman    time.Time
	Tgl_pengembalian  time.Time
	Durasi_Peminjaman string `json:"Durasi_Peminjaman" form:"Durasi_Peminjaman"`
	Created_at        time.Time
	Updated_at        time.Time
}

type User struct {
	Id            int64  `json:"Id" form:"Id"`
	Id_buku       int64  `json:"Id_buku" form:"Id_buku"`
	Nama_Peminjam string `json:"Nama_peminjaman" form:"Nama_peminjaman"`
	Program_studi string `json:"Program_studi" form:"Program_studi"`
	Angkatan      int    `json:"Angkatan" form:"Angkatan"`
	No_telp       int    `json:"No_telp" form:"No_telp"`
	Jenis_kelamin string `json:"Jenis_kelamin" form:"Jenis_kelamin"`
	Created_at    time.Time
	Updated_at    time.Time
}
