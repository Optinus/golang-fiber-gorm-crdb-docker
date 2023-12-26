package router

import (
	// Import paket-paket yang diperlukan
	"github.com/Optinus/golang-fiber-gorm-crdb-docker/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRouter mengatur rute-rute HTTP untuk aplikasi
func SetupRouter(app *fiber.App) {
	// Rute untuk membuat Todo baru menggunakan metode HTTP POST
	app.Post("/", handler.CreateTodo)

	// Rute untuk mendapatkan Todo berdasarkan ID menggunakan metode HTTP GET
	app.Get("/:id", handler.GetTodoById)

	// Rute untuk memperbarui Todo berdasarkan ID menggunakan metode HTTP PUT
	app.Put("/:id", handler.UpdateTodo)

	// Rute untuk menghapus Todo berdasarkan ID menggunakan metode HTTP DELETE
	app.Delete("/:id", handler.DeleteTodo)

	// Rute untuk mendapatkan semua Todo menggunakan metode HTTP GET
	app.Get("/", handler.GetAllTodo)
}
