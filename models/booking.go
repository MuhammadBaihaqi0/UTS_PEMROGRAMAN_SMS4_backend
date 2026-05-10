package models

import "time"

// Booking adalah struct utama untuk data booking lapangan
type Booking struct {
	ID            int       `json:"id"`
	NamaPemesan   string    `json:"nama_pemesan"`
	Email         string    `json:"email"`
	NoTelepon     string    `json:"no_telepon"`
	NamaLapangan  string    `json:"nama_lapangan"`
	JenisOlahraga string    `json:"jenis_olahraga"`
	TanggalMain   string    `json:"tanggal_main"`
	JamMulai      string    `json:"jam_mulai"`
	JamSelesai    string    `json:"jam_selesai"`
	TotalHarga    float64   `json:"total_harga"`
	Status        string    `json:"status"`
	Catatan       string    `json:"catatan"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// BookingRequest adalah struct untuk menerima input dari client
type BookingRequest struct {
	NamaPemesan   string  `json:"nama_pemesan"`
	Email         string  `json:"email"`
	NoTelepon     string  `json:"no_telepon"`
	NamaLapangan  string  `json:"nama_lapangan"`
	JenisOlahraga string  `json:"jenis_olahraga"`
	TanggalMain   string  `json:"tanggal_main"`
	JamMulai      string  `json:"jam_mulai"`
	JamSelesai    string  `json:"jam_selesai"`
	TotalHarga    float64 `json:"total_harga"`
	Status        string  `json:"status"`
	Catatan       string  `json:"catatan"`
}

// APIResponse adalah format response standar
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
