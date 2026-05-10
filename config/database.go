package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Gagal membuka koneksi database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Gagal terhubung ke database: %v", err)
	}

	DB = db
	log.Println("✅ Berhasil terhubung ke database Supabase PostgreSQL")
}
