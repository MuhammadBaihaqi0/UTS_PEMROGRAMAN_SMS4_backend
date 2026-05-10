package models

import "time"

// 1. FITUR LAPANGAN (Master Data)
type Lapangan struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	Nama          string      `json:"nama"`
	JenisOlahraga string      `json:"jenis_olahraga"` // Futsal, Basket, Badminton, dll
	HargaPerJam   float64     `json:"harga_per_jam"`
	Fasilitas     string      `json:"fasilitas"`
	GambarURL     string      `json:"gambar_url"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	// Relasi
	Pemesanan     []Pemesanan `gorm:"foreignKey:LapanganID" json:"pemesanan,omitempty"`
}

// 2. FITUR PEMESANAN (Transaksi Booking)
type Pemesanan struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	NamaPemesan  string     `json:"nama_pemesan"`
	Email        string     `json:"email"`
	NoTelepon    string     `json:"no_telepon"`
	
	// Relasi ke Lapangan
	LapanganID   uint       `json:"lapangan_id"`
	Lapangan     Lapangan   `gorm:"foreignKey:LapanganID" json:"lapangan"`
	
	// 3. FITUR JADWAL (Disimpan dalam transaksi pemesanan)
	TanggalMain  string     `json:"tanggal_main"` // Format: YYYY-MM-DD
	JamMulai     string     `json:"jam_mulai"`    // Format: HH:MM
	JamSelesai   string     `json:"jam_selesai"`  // Format: HH:MM
	
	TotalHarga   float64    `json:"total_harga"`
	Catatan      string     `json:"catatan"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	
	// Relasi ke Pembayaran
	Pembayaran   Pembayaran `gorm:"foreignKey:PemesananID" json:"pembayaran"`
}

// 4. FITUR PEMBAYARAN (Transaksi Keuangan)
type Pembayaran struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PemesananID   uint      `json:"pemesanan_id"`
	MetodeBayar   string    `json:"metode_bayar"` // Transfer Bank, E-Wallet, Tunai
	JumlahBayar   float64   `json:"jumlah_bayar"`
	StatusBayar   string    `json:"status_bayar"` // unpaid, paid, failed, refunded
	BuktiTransfer string    `json:"bukti_transfer"` // URL gambar bukti bayar
	WaktuBayar    time.Time `json:"waktu_bayar"`
}