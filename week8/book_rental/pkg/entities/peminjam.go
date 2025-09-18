package entities

import "time"

type Peminjam struct {
	ID                  int       `json:"ID"`
	Nama                string    `json:"Nama"`
	TanggalPeminjaman   time.Time `json:"TanggalPeminjaman"`
	TanggalPengembalian time.Time `json:"TanggalPengembalian"`
	BukuPinjaman        []Buku
}
