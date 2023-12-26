package handler

import (
	// Import paket-paket yang diperlukan
	database "github.com/Optinus/golang-fiber-gorm-crdb-docker/database"
	"github.com/Optinus/golang-fiber-gorm-crdb-docker/model"
	"github.com/gofiber/fiber/v2"
)

// CreateTodo membuat tugas baru berdasarkan data dari body permintaan
func CreateTodo(c *fiber.Ctx) error {
	// Membuat variabel baru untuk menyimpan data tugas baru
	var newTodo model.Todo

	// Mengurai data dari body permintaan dan memasukkannya ke dalam variabel newTodo
	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Menyimpan data tugas baru ke dalam database
	if err := database.DB.Create(&newTodo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create todo",
			"error":   err.Error(),
		})
	}

	// Mengembalikan data tugas baru dalam respons JSON
	return c.JSON(newTodo)
}

// GetTodoById mendapatkan tugas berdasarkan ID dari database
func GetTodoById(c *fiber.Ctx) error {
	// Mengambil ID dari parameter URL
	id := c.Params("id")
	var todo model.Todo

	// Mencari tugas berdasarkan ID dalam database
	if err := database.DB.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Todo not found",
			"error":   err.Error(),
		})
	}

	// Mengembalikan data tugas dalam respons JSON
	return c.JSON(todo)
}

// UpdateTodo memperbarui tugas berdasarkan ID dengan data dari body permintaan
func UpdateTodo(c *fiber.Ctx) error {
	// Mengambil ID dari parameter URL
	id := c.Params("id")
	var updateTodo model.Todo

	// Mencari tugas berdasarkan ID dalam database
	if err := database.DB.First(&updateTodo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Todo not found",
			"error":   err.Error(),
		})
	}

	// Mengurai data dari body permintaan dan memasukkannya ke dalam variabel updateTodo
	if err := c.BodyParser(&updateTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Menyimpan perubahan pada tugas ke dalam database
	if err := database.DB.Save(&updateTodo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update todo",
			"error":   err.Error(),
		})
	}

	// Mengembalikan data tugas yang telah diperbarui dalam respons JSON
	return c.JSON(updateTodo)
}

// DeleteTodo menghapus tugas berdasarkan ID dari database
func DeleteTodo(c *fiber.Ctx) error {
	// Mengambil ID dari parameter URL
	id := c.Params("id")
	var deleteTodo model.Todo

	// Mencari tugas berdasarkan ID dalam database
	if err := database.DB.First(&deleteTodo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Todo not found",
			"error":   err.Error(),
		})
	}

	// Menghapus tugas dari database
	if err := database.DB.Delete(&deleteTodo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete todo",
			"error":   err.Error(),
		})
	}

	// Mengembalikan respons bahwa tugas telah dihapus dengan sukses
	return c.JSON(fiber.Map{"message": "Todo deleted successfully"})
}

// GetAllTodo mendapatkan semua tugas dari database
func GetAllTodo(c *fiber.Ctx) error {
	// Membuat variabel untuk menyimpan semua tugas dari database
	var todo []model.Todo

	// Mengambil semua tugas dari database
	database.DB.Find(&todo)

	// Mengembalikan data semua tugas dalam respons JSON
	return c.JSON(todo)
}
