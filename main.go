package main

import (
	"booking-lapangan-backend/config"
	"booking-lapangan-backend/middleware"
	"booking-lapangan-backend/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  File .env tidak ditemukan, menggunakan environment variable sistem")
	}

	// Koneksi ke database
	config.InitDB()

	// Inisialisasi Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Booking Lapangan Olahraga API v1.0",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		},
	})

	// Setup middleware
	middleware.SetupCORS(app)
	middleware.SetupLogger(app)

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server berjalan di http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
