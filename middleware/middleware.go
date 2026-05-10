package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupCORS mengatur middleware CORS
func SetupCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
}

// SetupLogger mengatur middleware logger custom
func SetupLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} ${method} ${path} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
}

// RequestTimer adalah middleware custom untuk mencatat waktu request
func RequestTimer(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	log.Printf("📡 %s %s → %d [%v]", c.Method(), c.Path(), c.Response().StatusCode(), duration)
	return err
}
