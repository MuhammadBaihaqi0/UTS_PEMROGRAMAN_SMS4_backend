package routes

import (
	"booking-lapangan-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "🏟️ Booking Lapangan Olahraga API - Server berjalan dengan baik!",
			"version": "1.0.0",
		})
	})

	// API v1
	api := app.Group("/api/v1")

	// Booking routes
	bookings := api.Group("/bookings")
	bookings.Get("/", handlers.GetAllBookings)        // GET All
	bookings.Get("/:id", handlers.GetBookingByID)     // GET By ID
	bookings.Post("/", handlers.CreateBooking)        // POST Create
	bookings.Put("/:id", handlers.UpdateBooking)      // PUT Update
	bookings.Delete("/:id", handlers.DeleteBooking)   // DELETE

	// Lapangan routes
	lapangan := api.Group("/lapangan")
	lapangan.Get("/", handlers.GetAllLapangan)
	lapangan.Get("/:id", handlers.GetLapanganByID)
	lapangan.Post("/", handlers.CreateLapangan)
	lapangan.Put("/:id", handlers.UpdateLapangan)
	lapangan.Delete("/:id", handlers.DeleteLapangan)

	// Pemesanan routes
	pemesanan := api.Group("/pemesanan")
	pemesanan.Get("/", handlers.GetAllPemesanan)
	pemesanan.Get("/:id", handlers.GetPemesananByID)
	pemesanan.Post("/", handlers.CreatePemesanan)
	pemesanan.Put("/:id", handlers.UpdatePemesanan)
	pemesanan.Delete("/:id", handlers.DeletePemesanan)

	// Jadwal routes
	api.Get("/jadwal/:lapangan_id", handlers.GetJadwalByLapangan)

	// Pembayaran routes
	pembayaran := api.Group("/pembayaran")
	pembayaran.Get("/", handlers.GetAllPembayaran)
	pembayaran.Get("/:id", handlers.GetPembayaranByID)
	pembayaran.Post("/", handlers.CreatePembayaran)
	pembayaran.Put("/:id", handlers.UpdatePembayaran)
	pembayaran.Delete("/:id", handlers.DeletePembayaran)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Endpoint tidak ditemukan",
		})
	})
}
