package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			err = godotenv.Load("../.env")
		}
		if err != nil {
			log.Println(".env tidak ditemukan")
		}
	}

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan dalam .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database: %v", err)
	} 

	DB = db
	fmt.Println("Koneksi ke postgres berhasil")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum diinisialisasi00")
	}
	return DB
}
