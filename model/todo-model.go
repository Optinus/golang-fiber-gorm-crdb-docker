package model

// Todo adalah struktur data untuk merepresentasikan sebuah tugas (todo)
type Todo struct {
	ID          int    `json:"id"`          // ID adalah identifikasi unik untuk setiap tugas
	Title       string `json:"title"`       // Title adalah judul atau nama untuk tugas
	Description string `json:"description"` // Description adalah deskripsi atau keterangan untuk tugas
	Finished    bool   `json:"finished"`    // Finished menunjukkan apakah tugas sudah selesai atau belum
}
