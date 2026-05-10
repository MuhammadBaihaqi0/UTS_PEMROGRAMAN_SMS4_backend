//go:build ignore

package main


import (
	"booking-lapangan-backend/config"
	"booking-lapangan-backend/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.InitDB()

	var count int64
	config.DB.Model(&models.Booking{}).Count(&count)

	if count >= 10 {
		fmt.Printf("Database sudah memiliki %d data, tidak perlu seeder.\n", count)
		return
	}

	dummyBookings := []models.Booking{
		{NamaPemesan: "Budi Santoso", Email: "budi@example.com", NoTelepon: "081234567890", NamaLapangan: "Arena Futsal Jakarta", JenisOlahraga: "Futsal", TanggalMain: "2026-06-01", JamMulai: "10:00", JamSelesai: "12:00", TotalHarga: 300000, Catatan: "Tolong siapkan bola"},
		{NamaPemesan: "Andi Wijaya", Email: "andi@example.com", NoTelepon: "081234567891", NamaLapangan: "Gor Basket Bandung", JenisOlahraga: "Basket", TanggalMain: "2026-06-02", JamMulai: "14:00", JamSelesai: "16:00", TotalHarga: 400000, Catatan: ""},
		{NamaPemesan: "Siti Aminah", Email: "siti@example.com", NoTelepon: "081234567892", NamaLapangan: "Badminton Center SBY", JenisOlahraga: "Badminton", TanggalMain: "2026-06-03", JamMulai: "08:00", JamSelesai: "10:00", TotalHarga: 200000, Catatan: ""},
		{NamaPemesan: "Rudi Hartono", Email: "rudi@example.com", NoTelepon: "081234567893", NamaLapangan: "Tenis Court Bali", JenisOlahraga: "Tenis", TanggalMain: "2026-06-04", JamMulai: "16:00", JamSelesai: "18:00", TotalHarga: 500000, Catatan: ""},
		{NamaPemesan: "Dina Fitriani", Email: "dina@example.com", NoTelepon: "081234567894", NamaLapangan: "Voli Arena JKT", JenisOlahraga: "Voli", TanggalMain: "2026-06-05", JamMulai: "19:00", JamSelesai: "21:00", TotalHarga: 300000, Catatan: "Net voli minta yg baru"},
		{NamaPemesan: "Eko Prasetyo", Email: "eko@example.com", NoTelepon: "081234567895", NamaLapangan: "Arena Futsal Jakarta", JenisOlahraga: "Futsal", TanggalMain: "2026-06-06", JamMulai: "20:00", JamSelesai: "22:00", TotalHarga: 300000, Catatan: ""},
		{NamaPemesan: "Rini Astuti", Email: "rini@example.com", NoTelepon: "081234567896", NamaLapangan: "Gor Basket Bandung", JenisOlahraga: "Basket", TanggalMain: "2026-06-07", JamMulai: "09:00", JamSelesai: "11:00", TotalHarga: 400000, Catatan: ""},
		{NamaPemesan: "Galih Pratama", Email: "galih@example.com", NoTelepon: "081234567897", NamaLapangan: "Badminton Center SBY", JenisOlahraga: "Badminton", TanggalMain: "2026-06-08", JamMulai: "15:00", JamSelesai: "17:00", TotalHarga: 200000, Catatan: ""},
		{NamaPemesan: "Putri Rahayu", Email: "putri@example.com", NoTelepon: "081234567898", NamaLapangan: "Tenis Court Bali", JenisOlahraga: "Tenis", TanggalMain: "2026-06-09", JamMulai: "07:00", JamSelesai: "09:00", TotalHarga: 500000, Catatan: ""},
		{NamaPemesan: "Hendra Gunawan", Email: "hendra@example.com", NoTelepon: "081234567899", NamaLapangan: "Arena Futsal Jakarta", JenisOlahraga: "Futsal", TanggalMain: "2026-06-10", JamMulai: "18:00", JamSelesai: "20:00", TotalHarga: 300000, Catatan: ""},
	}

	for _, b := range dummyBookings {
		config.DB.Create(&b)
	}
	fmt.Printf("Berhasil melakukan seeding %d data booking!\n", len(dummyBookings))
}
