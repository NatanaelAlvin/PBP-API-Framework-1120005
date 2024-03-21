package controller

import (
	"database/sql"
	"log"

	//"gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	//"gorm.io/gorm"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp")
	if err != nil {
		log.Println("Error executing query:", err)
	}
	return db
}

//var db *gorm.DB

// InitDB inisialisasi koneksi database
//
