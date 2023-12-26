package database

import (
	// Import paket-paket yang diperlukan

	"github.com/Optinus/golang-fiber-gorm-crdb/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah objek GORM untuk berinteraksi dengan database
var DB *gorm.DB

// ConnectDatabase menghubungkan aplikasi ke database dan melakukan migrasi otomatis untuk model Todo
func ConnectDatabase() {
	// Membuka koneksi ke database CockroachDB menggunakan driver PostgreSQL
	db, err := gorm.Open(postgres.Open("postgresql://root@192.168.0.104:26257?sslmode=disable"))
	if err != nil {
		// Panic jika terjadi kesalahan dalam membuka koneksi
		panic(err)
	}

	// Melakukan migrasi otomatis untuk model Todo
	db.AutoMigrate(&model.Todo{})

	// Menyimpan objek GORM ke dalam variabel global DB
	DB = db
}
