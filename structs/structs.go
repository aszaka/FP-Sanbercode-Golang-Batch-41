package structs

import (
	"cloud.google.com/go/civil"
)

type User struct {
	IdUser   int16  `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Level    int8   `json:"level"`
}

type MasterPS struct {
	IdPS             int16  `json:"id_ps"`
	TipePS           string `json:"tipe_ps"`
	DaftarGame       string `json:"daftar_game"`
	StatusPeminjaman int8   `json:"status_peminjaman"`
}

type Harga struct {
	IdHarga int16  `json:"id_harga"`
	Waktu   int16  `json:"waktu"`
	Satuan  string `json:"satuan"`
	Harga   int16  `json:"harga"`
}

type Peminjaman struct {
	IdPeminjaman int16          `json:"id_peminjaman"`
	IdPS         int16          `json:"id_ps"`
	IdUser       int16          `json:"id_user"`
	IdHarga      int16          `json:"id_harga"`
	Peminjaman   civil.DateTime `json:"peminjaman"`
	Pengembalian civil.DateTime `json:"pengembalian"`
	Keterangan   string         `json:"keterangan"`
}
