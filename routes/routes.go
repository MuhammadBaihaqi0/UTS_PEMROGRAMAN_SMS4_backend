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

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Endpoint tidak ditemukan",
		})
	})
}
