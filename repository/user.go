package repository

import (
	"FP/structs"
	"database/sql"
)

func GetAllUser(db *sql.DB) (err error, results []structs.User) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.IdUser, &user.Username, &user.Password, &user.Level)
		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	return
}

func AddUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (username,password,level) VALUES ($1,$2,2)"

	errs := db.QueryRow(sql, user.Username, user.Password)

	return errs.Err()
}

func AddAdmin(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (username,password,level) VALUES ($1,$2,1)"

	errs := db.QueryRow(sql, user.Username, user.Password)

	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET username = $1, password = $2, level = $3 WHERE id_user= $4"

	errs := db.QueryRow(sql, user.Username, user.Password, user.Level, user.IdUser)

	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sql := "DELETE FROM users WHERE id_user = $1"

	errs := db.QueryRow(sql, user.IdUser)

	return errs.Err()
}
