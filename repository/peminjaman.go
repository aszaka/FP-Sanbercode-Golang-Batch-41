package repository

import (
	"FP/structs"
	"database/sql"
)

func GetAllPeminjaman(db *sql.DB) (err error, results []structs.Peminjaman) {
	sql := "SELECT * FROM peminjaman"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var peminjaman = structs.Peminjaman{}

		err = rows.Scan(&peminjaman.IdPeminjaman,
			&peminjaman.IdPS,
			&peminjaman.IdUser,
			&peminjaman.IdHarga,
			&peminjaman.Peminjaman,
			&peminjaman.Pengembalian,
			&peminjaman.Keterangan)
		if err != nil {
			panic(err)
		}

		results = append(results, peminjaman)
	}

	return
}

func AddPeminjaman(db *sql.DB, peminjaman structs.Peminjaman) (err error) {
	sql := "INSERT INTO peminjaman (id_ps, id_user, id_harga, peminjaman, pengembalian, keterangan) VALUES ($1,$2,$3,$4,$5,$6)"

	errs := db.QueryRow(sql, peminjaman.IdPS, peminjaman.IdUser, peminjaman.IdHarga, peminjaman.Peminjaman, peminjaman.Pengembalian, peminjaman.Keterangan)

	return errs.Err()
}

func UpdatePeminjaman(db *sql.DB, peminjaman structs.Peminjaman) (err error) {
	sql := "UPDATE peminjaman SET id_ps = $1, id_user = $2, id_harga = $3, peminjaman=$4, pengembalian=$5, keterangan=$6 WHERE id_peminjaman= $7"

	errs := db.QueryRow(sql, peminjaman.IdPS, peminjaman.IdUser, peminjaman.IdHarga, peminjaman.Peminjaman, peminjaman.Pengembalian, peminjaman.Keterangan, peminjaman.IdPeminjaman)

	return errs.Err()
}

func DeletePeminjaman(db *sql.DB, peminjaman structs.Peminjaman) (err error) {
	sql := "DELETE FROM peminjaman WHERE id_peminjaman = $1"

	errs := db.QueryRow(sql, peminjaman.IdPeminjaman)

	return errs.Err()
}
