package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := "user=username dbname=catfoodstore sslmode=disable password=yourpassword"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		println("ошибка подключения к БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		println("не удалось подключиться к базе данных: %v", err)
	}
	DB = db
	println("Успешно подключились к базе данных")

	return nil
}
