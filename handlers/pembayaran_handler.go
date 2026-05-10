package handlers

import (
	"booking-lapangan-backend/config"
	"booking-lapangan-backend/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetAllPembayaran
func GetAllPembayaran(c *fiber.Ctx) error {
	var pembayaran []models.Pembayaran
	result := config.DB.Find(&pembayaran)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengambil data pembayaran",
			Error:   result.Error.Error(),
		})
	}
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data pembayaran",
		Data:    pembayaran,
	})
}

// GetPembayaranByID
func GetPembayaranByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pembayaran
	result := config.DB.First(&p, id)
	if result.Error != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Pembayaran dengan ID %s tidak ditemukan", id),
		})
	}
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data pembayaran",
		Data:    p,
	})
}

// CreatePembayaran
func CreatePembayaran(c *fiber.Ctx) error {
	var p models.Pembayaran
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
			Error:   err.Error(),
		})
	}

	p.WaktuBayar = time.Now()
	p.StatusBayar = "paid"

	result := config.DB.Create(&p)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal membuat pembayaran",
			Error:   result.Error.Error(),
		})
	}
	
	// Update status pemesanan
	var pemesanan models.Pemesanan
	if err := config.DB.First(&pemesanan, p.PemesananID).Error; err == nil {
		pemesanan.Status = "confirmed"
		config.DB.Save(&pemesanan)
	}

	return c.Status(201).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil memproses pembayaran",
		Data:    p,
	})
}

// UpdatePembayaran
func UpdatePembayaran(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pembayaran
	if err := config.DB.First(&p, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Pembayaran tidak ditemukan",
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
		Message: "Berhasil update pembayaran",
		Data:    p,
	})
}

// DeletePembayaran
func DeletePembayaran(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Pembayaran
	if err := config.DB.First(&p, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Pembayaran tidak ditemukan",
		})
	}

	config.DB.Delete(&p)
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil menghapus pembayaran",
	})
}
