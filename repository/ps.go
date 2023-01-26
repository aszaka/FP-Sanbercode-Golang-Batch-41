package repository

import (
	"FP/structs"
	"database/sql"
)

func GetAllPS(db *sql.DB) (err error, results []structs.MasterPS) {
	sql := "SELECT * FROM master_ps"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var ps = structs.MasterPS{}

		err = rows.Scan(&ps.IdPS, &ps.TipePS, &ps.DaftarGame, &ps.StatusPeminjaman)
		if err != nil {
			panic(err)
		}

		results = append(results, ps)
	}

	return
}

func AddPS(db *sql.DB, ps structs.MasterPS) (err error) {
	sql := "INSERT INTO master_ps (tipe_ps,daftar_game,status_peminjaman) VALUES ($1,$2,0)"

	errs := db.QueryRow(sql, ps.TipePS, ps.DaftarGame)

	return errs.Err()
}

func UpdatePS(db *sql.DB, ps structs.MasterPS) (err error) {
	sql := "UPDATE master_ps SET tipe_ps = $1, daftar_game = $2, status_peminjaman = $3 WHERE id_ps= $4"

	errs := db.QueryRow(sql, ps.TipePS, ps.DaftarGame, ps.StatusPeminjaman, ps.IdPS)

	return errs.Err()
}

func DeletePS(db *sql.DB, ps structs.MasterPS) (err error) {
	sql := "DELETE FROM master_ps WHERE id_ps = $1"

	errs := db.QueryRow(sql, ps.IdPS)

	return errs.Err()
}
