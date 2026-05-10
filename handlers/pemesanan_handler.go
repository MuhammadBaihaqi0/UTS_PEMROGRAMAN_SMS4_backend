package handlers

import (
	"booking-lapangan-backend/config"
	"booking-lapangan-backend/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllPemesanan
func GetAllPemesanan(c *fiber.Ctx) error {
	var pemesanan []models.Pemesanan
	result := config.DB.Preload("Lapangan").Preload("Pembayaran").Find(&pemesanan)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengambil data pemesanan",
			Error:   result.Error.Error(),
		})
	}
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data pemesanan",
		Data:    pemesanan,
	})
}

// GetPemesananByID
func GetPemesananByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pemesanan
	result := config.DB.Preload("Lapangan").Preload("Pembayaran").First(&p, id)
	if result.Error != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Pemesanan dengan ID %s tidak ditemukan", id),
		})
	}
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data pemesanan",
		Data:    p,
	})
}

// CreatePemesanan
func CreatePemesanan(c *fiber.Ctx) error {
	var p models.Pemesanan
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
			Error:   err.Error(),
		})
	}

	p.Status = "pending"

	result := config.DB.Create(&p)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal membuat pemesanan",
			Error:   result.Error.Error(),
		})
	}

	return c.Status(201).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil membuat pemesanan",
		Data:    p,
	})
}

// UpdatePemesanan
func UpdatePemesanan(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pemesanan
	if err := config.DB.First(&p, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Pemesanan tidak ditemukan",
		})
	}

	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
		})
	}

	config.DB.Save(&p)
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil update pemesanan",
		Data:    p,
	})
}

// DeletePemesanan
func DeletePemesanan(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pemesanan
	if err := config.DB.First(&p, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Pemesanan tidak ditemukan",
		})
	}

	config.DB.Delete(&p)
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil menghapus pemesanan",
	})
}

// GetJadwalByLapangan
func GetJadwalByLapangan(c *fiber.Ctx) error {
	lapanganID := c.Params("lapangan_id")
	var pemesanan []models.Pemesanan
	result := config.DB.Where("lapangan_id = ? AND status != 'cancelled'", lapanganID).Find(&pemesanan)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengambil jadwal",
			Error:   result.Error.Error(),
		})
	}
	
	// Simplify payload for jadwal
	var jadwal []map[string]interface{}
	for _, p := range pemesanan {
		jadwal = append(jadwal, map[string]interface{}{
			"id": p.ID,
			"tanggal_main": p.TanggalMain,
			"jam_mulai": p.JamMulai,
			"jam_selesai": p.JamSelesai,
			"status": p.Status,
		})
	}

	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil jadwal",
		Data:    jadwal,
	})
}
