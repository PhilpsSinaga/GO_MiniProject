package model

import (
	"time"
)

type Buku struct {
	Id                  int64
	Nama_buku           string
	Penulis             string
	Penerbit            string
	Tahun_buku          int
	Edisi               string
	Jumlah_halaman      int
	Jumlah_halaman_baca int
	Jenis_buku          string
	Jumlah_buku         int
	Ketersedian_buku    string
	Created_at          time.Time
	Updated_at          time.Time
}

type Lokasi_buku struct {
	Id         int64
	Id_buku    int64
	Lantai     int
	Rak_buku   int
	baris_buku int
	Created_at time.Time
	Updated_at time.Time
}

type Peminjamans struct {
	Id                int64
	Id_buku           int64
	Tgl_peminjaman    time.Time
	Tgl_pengembalian  time.Time
	Durasi_Peminjaman string
	Created_at        time.Time
	Updated_at        time.Time
}

type Peminjams struct {
	Id            int64
	Id_buku       int64
	Nama_Peminjam string
	Program_studi string
	Angkatan      int
	No_telp       int
	Jenis_kelamin string
	Created_at    time.Time
	Updated_at    time.Time
}

type Petugas struct {
	Id         int64
	Username   string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
