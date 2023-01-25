package structs

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
