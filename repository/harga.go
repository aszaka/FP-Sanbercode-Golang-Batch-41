package repository

import (
	"FP/structs"
	"database/sql"
)

func GetAllHarga(db *sql.DB) (err error, results []structs.Harga) {
	sql := "SELECT * FROM harga"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var harga = structs.Harga{}

		err = rows.Scan(&harga.IdHarga, &harga.Waktu, &harga.Satuan, &harga.Harga)
		if err != nil {
			panic(err)
		}

		results = append(results, harga)
	}

	return
}

func AddHarga(db *sql.DB, harga structs.Harga) (err error) {
	sql := "INSERT INTO harga (waktu, satuan, harga) VALUES ($1,$2,$3)"

	errs := db.QueryRow(sql, harga.Waktu, harga.Satuan, harga.Harga)

	return errs.Err()
}

func UpdateHarga(db *sql.DB, harga structs.Harga) (err error) {
	sql := "UPDATE harga SET waktu = $1, satuan = $2, harga = $3 WHERE id_harga= $4"

	errs := db.QueryRow(sql, harga.Waktu, harga.Satuan, harga.Harga, harga.IdHarga)

	return errs.Err()
}

func DeleteHarga(db *sql.DB, harga structs.Harga) (err error) {
	sql := "DELETE FROM harga WHERE id_harga = $1"

	errs := db.QueryRow(sql, harga.IdHarga)

	return errs.Err()
}
