package handlers

import (
	"booking-lapangan-backend/config"
	"booking-lapangan-backend/models"
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ─── HELPER: Validasi Input ───────────────────────────────────────────────────

func validateBookingRequest(req models.BookingRequest) []string {
	var errors []string

	// Nama pemesan: wajib, min 3, max 100 karakter
	req.NamaPemesan = strings.TrimSpace(req.NamaPemesan)
	if req.NamaPemesan == "" {
		errors = append(errors, "nama_pemesan wajib diisi")
	} else if len(req.NamaPemesan) < 3 {
		errors = append(errors, "nama_pemesan minimal 3 karakter")
	} else if len(req.NamaPemesan) > 100 {
		errors = append(errors, "nama_pemesan maksimal 100 karakter")
	}

	// Email: wajib, format valid
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		errors = append(errors, "email wajib diisi")
	} else {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(req.Email) {
			errors = append(errors, "format email tidak valid")
		}
	}

	// No Telepon: wajib, min 10, max 15 karakter
	req.NoTelepon = strings.TrimSpace(req.NoTelepon)
	if req.NoTelepon == "" {
		errors = append(errors, "no_telepon wajib diisi")
	} else if len(req.NoTelepon) < 10 {
		errors = append(errors, "no_telepon minimal 10 karakter")
	} else if len(req.NoTelepon) > 15 {
		errors = append(errors, "no_telepon maksimal 15 karakter")
	}

	// Nama Lapangan: wajib
	if strings.TrimSpace(req.NamaLapangan) == "" {
		errors = append(errors, "nama_lapangan wajib diisi")
	}

	// Jenis Olahraga: wajib
	if strings.TrimSpace(req.JenisOlahraga) == "" {
		errors = append(errors, "jenis_olahraga wajib diisi")
	}

	// Tanggal Main: wajib
	if strings.TrimSpace(req.TanggalMain) == "" {
		errors = append(errors, "tanggal_main wajib diisi")
	}

	// Jam Mulai & Selesai: wajib
	if strings.TrimSpace(req.JamMulai) == "" {
		errors = append(errors, "jam_mulai wajib diisi")
	}
	if strings.TrimSpace(req.JamSelesai) == "" {
		errors = append(errors, "jam_selesai wajib diisi")
	}

	// Total Harga: tidak boleh negatif
	if req.TotalHarga < 0 {
		errors = append(errors, "total_harga tidak boleh negatif")
	}

	// Status: wajib & harus salah satu dari nilai yang valid
	validStatus := map[string]bool{"pending": true, "confirmed": true, "cancelled": true, "completed": true}
	if req.Status == "" {
		errors = append(errors, "status wajib diisi")
	} else if !validStatus[req.Status] {
		errors = append(errors, "status harus salah satu dari: pending, confirmed, cancelled, completed")
	}

	return errors
}

// ─── GET ALL BOOKINGS ─────────────────────────────────────────────────────────

func GetAllBookings(c *fiber.Ctx) error {
	rows, err := config.DB.Query(`
		SELECT id, nama_pemesan, email, no_telepon, nama_lapangan, jenis_olahraga,
		       tanggal_main, jam_mulai, jam_selesai, total_harga, status, catatan,
		       created_at, updated_at
		FROM bookings
		ORDER BY created_at DESC
	`)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengambil data booking",
			Error:   err.Error(),
		})
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		err := rows.Scan(
			&b.ID, &b.NamaPemesan, &b.Email, &b.NoTelepon,
			&b.NamaLapangan, &b.JenisOlahraga, &b.TanggalMain,
			&b.JamMulai, &b.JamSelesai, &b.TotalHarga,
			&b.Status, &b.Catatan, &b.CreatedAt, &b.UpdatedAt,
		)
		if err != nil {
			return c.Status(500).JSON(models.APIResponse{
				Success: false,
				Message: "Gagal membaca data booking",
				Error:   err.Error(),
			})
		}
		bookings = append(bookings, b)
	}

	if bookings == nil {
		bookings = []models.Booking{}
	}

	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Berhasil mengambil %d data booking", len(bookings)),
		Data:    bookings,
	})
}

// ─── GET BOOKING BY ID ────────────────────────────────────────────────────────

func GetBookingByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var b models.Booking
	err := config.DB.QueryRow(`
		SELECT id, nama_pemesan, email, no_telepon, nama_lapangan, jenis_olahraga,
		       tanggal_main, jam_mulai, jam_selesai, total_harga, status, catatan,
		       created_at, updated_at
		FROM bookings WHERE id = $1
	`, id).Scan(
		&b.ID, &b.NamaPemesan, &b.Email, &b.NoTelepon,
		&b.NamaLapangan, &b.JenisOlahraga, &b.TanggalMain,
		&b.JamMulai, &b.JamSelesai, &b.TotalHarga,
		&b.Status, &b.Catatan, &b.CreatedAt, &b.UpdatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Booking dengan ID %s tidak ditemukan", id),
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data booking",
		Data:    b,
	})
}

// ─── CREATE BOOKING ───────────────────────────────────────────────────────────

func CreateBooking(c *fiber.Ctx) error {
	var req models.BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
			Error:   err.Error(),
		})
	}

	// Validasi input
	validationErrors := validateBookingRequest(req)
	if len(validationErrors) > 0 {
		return c.Status(422).JSON(fiber.Map{
			"success":  false,
			"message":  "Validasi gagal",
			"errors":   validationErrors,
		})
	}

	// Cek duplikat: email + lapangan + tanggal + jam_mulai
	var count int
	err := config.DB.QueryRow(`
		SELECT COUNT(*) FROM bookings
		WHERE nama_lapangan = $1 AND tanggal_main = $2 AND jam_mulai = $3
	`, req.NamaLapangan, req.TanggalMain, req.JamMulai).Scan(&count)
	if err == nil && count > 0 {
		return c.Status(409).JSON(models.APIResponse{
			Success: false,
			Message: "Lapangan sudah dibooking pada waktu tersebut",
		})
	}

	var id int
	err = config.DB.QueryRow(`
		INSERT INTO bookings 
		(nama_pemesan, email, no_telepon, nama_lapangan, jenis_olahraga, tanggal_main, jam_mulai, jam_selesai, total_harga, status, catatan)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`,
		strings.TrimSpace(req.NamaPemesan),
		strings.TrimSpace(req.Email),
		strings.TrimSpace(req.NoTelepon),
		strings.TrimSpace(req.NamaLapangan),
		strings.TrimSpace(req.JenisOlahraga),
		req.TanggalMain,
		req.JamMulai,
		req.JamSelesai,
		req.TotalHarga,
		req.Status,
		strings.TrimSpace(req.Catatan),
	).Scan(&id)

	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal menyimpan data booking",
			Error:   err.Error(),
		})
	}

	// Ambil data yang baru dibuat
	var b models.Booking
	config.DB.QueryRow(`
		SELECT id, nama_pemesan, email, no_telepon, nama_lapangan, jenis_olahraga,
		       tanggal_main, jam_mulai, jam_selesai, total_harga, status, catatan,
		       created_at, updated_at
		FROM bookings WHERE id = $1
	`, id).Scan(
		&b.ID, &b.NamaPemesan, &b.Email, &b.NoTelepon,
		&b.NamaLapangan, &b.JenisOlahraga, &b.TanggalMain,
		&b.JamMulai, &b.JamSelesai, &b.TotalHarga,
		&b.Status, &b.Catatan, &b.CreatedAt, &b.UpdatedAt,
	)

	return c.Status(201).JSON(models.APIResponse{
		Success: true,
		Message: "Booking berhasil dibuat",
		Data:    b,
	})
}

// ─── UPDATE BOOKING ───────────────────────────────────────────────────────────

func UpdateBooking(c *fiber.Ctx) error {
	id := c.Params("id")

	// Cek apakah data ada
	var exists int
	config.DB.QueryRow("SELECT COUNT(*) FROM bookings WHERE id = $1", id).Scan(&exists)
	if exists == 0 {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Booking dengan ID %s tidak ditemukan", id),
		})
	}

	var req models.BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
			Error:   err.Error(),
		})
	}

	// Validasi input
	validationErrors := validateBookingRequest(req)
	if len(validationErrors) > 0 {
		return c.Status(422).JSON(fiber.Map{
			"success": false,
			"message": "Validasi gagal",
			"errors":  validationErrors,
		})
	}

	_, err := config.DB.Exec(`
		UPDATE bookings SET
			nama_pemesan = $1, email = $2, no_telepon = $3,
			nama_lapangan = $4, jenis_olahraga = $5, tanggal_main = $6,
			jam_mulai = $7, jam_selesai = $8, total_harga = $9,
			status = $10, catatan = $11, updated_at = NOW()
		WHERE id = $12
	`,
		strings.TrimSpace(req.NamaPemesan),
		strings.TrimSpace(req.Email),
		strings.TrimSpace(req.NoTelepon),
		strings.TrimSpace(req.NamaLapangan),
		strings.TrimSpace(req.JenisOlahraga),
		req.TanggalMain,
		req.JamMulai,
		req.JamSelesai,
		req.TotalHarga,
		req.Status,
		strings.TrimSpace(req.Catatan),
		id,
	)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengupdate data booking",
			Error:   err.Error(),
		})
	}

	// Ambil data terbaru
	var b models.Booking
	config.DB.QueryRow(`
		SELECT id, nama_pemesan, email, no_telepon, nama_lapangan, jenis_olahraga,
		       tanggal_main, jam_mulai, jam_selesai, total_harga, status, catatan,
		       created_at, updated_at
		FROM bookings WHERE id = $1
	`, id).Scan(
		&b.ID, &b.NamaPemesan, &b.Email, &b.NoTelepon,
		&b.NamaLapangan, &b.JenisOlahraga, &b.TanggalMain,
		&b.JamMulai, &b.JamSelesai, &b.TotalHarga,
		&b.Status, &b.Catatan, &b.CreatedAt, &b.UpdatedAt,
	)

	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Booking berhasil diupdate",
		Data:    b,
	})
}

// ─── DELETE BOOKING ───────────────────────────────────────────────────────────

func DeleteBooking(c *fiber.Ctx) error {
	id := c.Params("id")

	// Cek apakah data ada
	var exists int
	config.DB.QueryRow("SELECT COUNT(*) FROM bookings WHERE id = $1", id).Scan(&exists)
	if exists == 0 {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Booking dengan ID %s tidak ditemukan", id),
		})
	}

	_, err := config.DB.Exec("DELETE FROM bookings WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal menghapus data booking",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Booking dengan ID %s berhasil dihapus", id),
	})
}
