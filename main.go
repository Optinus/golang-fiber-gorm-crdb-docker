package main

import (
	// Import paket-paket yang diperlukan
	database "github.com/Optinus/golang-fiber-gorm-crdb-docker/database"
	"github.com/Optinus/golang-fiber-gorm-crdb-docker/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Hubungkan ke database
	database.ConnectDatabase()

	// Setup router dengan menggunakan fungsi dari paket router
	router.SetupRouter(app)

	// Mulai mendengarkan pada port 3000
	app.Listen(":3000")
}
