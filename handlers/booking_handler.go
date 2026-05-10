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

	return errors
}

// ─── GET ALL BOOKINGS ─────────────────────────────────────────────────────────

func GetAllBookings(c *fiber.Ctx) error {
	var bookings []models.Booking
	result := config.DB.Order("created_at DESC").Find(&bookings)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengambil data booking",
			Error:   result.Error.Error(),
		})
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
	result := config.DB.First(&b, id)

	if result.Error != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Booking dengan ID %s tidak ditemukan", id),
			Error:   result.Error.Error(),
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

	// Cek duplikat: lapangan + tanggal + jam_mulai
	var count int64
	config.DB.Model(&models.Booking{}).
		Where("nama_lapangan = ? AND tanggal_main = ? AND jam_mulai = ?",
			req.NamaLapangan, req.TanggalMain, req.JamMulai).
		Count(&count)
	if count > 0 {
		return c.Status(409).JSON(models.APIResponse{
			Success: false,
			Message: "Lapangan sudah dibooking pada waktu tersebut",
		})
	}

	// Buat booking baru
	b := models.Booking{
		NamaPemesan:   strings.TrimSpace(req.NamaPemesan),
		Email:         strings.TrimSpace(req.Email),
		NoTelepon:     strings.TrimSpace(req.NoTelepon),
		NamaLapangan:  strings.TrimSpace(req.NamaLapangan),
		JenisOlahraga: strings.TrimSpace(req.JenisOlahraga),
		TanggalMain:   req.TanggalMain,
		JamMulai:      req.JamMulai,
		JamSelesai:    req.JamSelesai,
		TotalHarga:    req.TotalHarga,
		Catatan:       strings.TrimSpace(req.Catatan),
	}

	result := config.DB.Create(&b)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal menyimpan data booking",
			Error:   result.Error.Error(),
		})
	}

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
	var b models.Booking
	result := config.DB.First(&b, id)
	if result.Error != nil {
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

	// Update fields
	b.NamaPemesan = strings.TrimSpace(req.NamaPemesan)
	b.Email = strings.TrimSpace(req.Email)
	b.NoTelepon = strings.TrimSpace(req.NoTelepon)
	b.NamaLapangan = strings.TrimSpace(req.NamaLapangan)
	b.JenisOlahraga = strings.TrimSpace(req.JenisOlahraga)
	b.TanggalMain = req.TanggalMain
	b.JamMulai = req.JamMulai
	b.JamSelesai = req.JamSelesai
	b.TotalHarga = req.TotalHarga
	b.Catatan = strings.TrimSpace(req.Catatan)

	result = config.DB.Save(&b)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengupdate data booking",
			Error:   result.Error.Error(),
		})
	}

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
	var b models.Booking
	result := config.DB.First(&b, id)
	if result.Error != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Booking dengan ID %s tidak ditemukan", id),
		})
	}

	result = config.DB.Delete(&b)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal menghapus data booking",
			Error:   result.Error.Error(),
		})
	}

	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Booking dengan ID %s berhasil dihapus", id),
	})
}
