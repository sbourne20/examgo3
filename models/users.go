package models

import (
	"database/sql"
	"fmt"
	"time"
)

type dateType time.Time

type User struct {
	Idpengguna    int    `json:"Id_pengguna"`
	Namapengguna  string `json:"Nama_pengguna"`
	Nomorrekening string `json:"Nomor_rekening"`
}

func GetUsers(db *sql.DB) ([]User, error) {
	statement := fmt.Sprintf("SELECT Id_pengguna, Nama_pengguna, Nomor_rekening FROM user_test")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User

		err = rows.Scan(
			&u.Idpengguna,
			&u.Namapengguna,
			&u.Nomorrekening)

		if err != nil {
			panic(err.Error())
		}
		users = append(users, u)
	}

	return users, nil
}

func (u *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO user_test(Id_pengguna, Nama_pengguna, Nomor_rekening) VALUES(%d, '%s', '%s')", u.Idpengguna, u.Namapengguna, u.Nomorrekening)

	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.Idpengguna)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE user_test SET Nama_pengguna='%s', Nomor_rekening='%s' WHERE Id_pengguna=%d", u.Namapengguna, u.Nomorrekening, u.Idpengguna)
	_, err := db.Exec(statement)
	return err
}

func (u *User) DeleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM user_test WHERE Id_pengguna=%d", u.Idpengguna)
	_, err := db.Exec(statement)
	return err
}
