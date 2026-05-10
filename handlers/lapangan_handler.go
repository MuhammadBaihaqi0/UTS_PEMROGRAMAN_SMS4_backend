package handlers

import (
	"booking-lapangan-backend/config"
	"booking-lapangan-backend/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllLapangan
func GetAllLapangan(c *fiber.Ctx) error {
	var lapangan []models.Lapangan
	result := config.DB.Find(&lapangan)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal mengambil data lapangan",
			Error:   result.Error.Error(),
		})
	}
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data lapangan",
		Data:    lapangan,
	})
}

// GetLapanganByID
func GetLapanganByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var l models.Lapangan
	result := config.DB.First(&l, id)
	if result.Error != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Lapangan dengan ID %s tidak ditemukan", id),
		})
	}
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil mengambil data lapangan",
		Data:    l,
	})
}

// CreateLapangan
func CreateLapangan(c *fiber.Ctx) error {
	var l models.Lapangan
	if err := c.BodyParser(&l); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
			Error:   err.Error(),
		})
	}

	result := config.DB.Create(&l)
	if result.Error != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Gagal membuat lapangan",
			Error:   result.Error.Error(),
		})
	}

	return c.Status(201).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil membuat lapangan",
		Data:    l,
	})
}

// UpdateLapangan
func UpdateLapangan(c *fiber.Ctx) error {
	id := c.Params("id")
	var l models.Lapangan
	if err := config.DB.First(&l, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Lapangan tidak ditemukan",
		})
	}

	if err := c.BodyParser(&l); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Format request tidak valid",
		})
	}

	config.DB.Save(&l)
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil update lapangan",
		Data:    l,
	})
}

// DeleteLapangan
func DeleteLapangan(c *fiber.Ctx) error {
	id := c.Params("id")
	var l models.Lapangan
	if err := config.DB.First(&l, id).Error; err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Lapangan tidak ditemukan",
		})
	}

	config.DB.Delete(&l)
	return c.Status(200).JSON(models.APIResponse{
		Success: true,
		Message: "Berhasil menghapus lapangan",
	})
}
