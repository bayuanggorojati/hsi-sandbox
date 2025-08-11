package main

import (
	"fmt"
	"time"
)

func cetakPesan(pesanan string, channelGoRoutine chan bool) {
	for i := 0; i < 5; i++ {
		println("Pesan ke-", i+1, ":", pesanan)
		time.Sleep(1 * time.Millisecond)
	}
	channelGoRoutine <- true // Mengirim data ke channel dengan nilai true
}

func main() {
	fmt.Println("Memulai Pesanan")

	//inisialisasi channel
	channel := make(chan bool, 2)

	// goRoutine Utama
	go cetakPesan("Nasi Goreng", channel)
	// goRoutine kedua
	go cetakPesan("Telur Mata Sapi", channel)

	// Menunggu goRoutine selesai
	// Dengan cara menunggu beberapa waktu
	time.Sleep(6 * time.Millisecond)
	// Tidak akan mencetak pesan "Selesai" sebelum channel menerima data dan bernilai true
	fmt.Println("Selesai")
}

// Pesanan 1 -> Waktu Proses
// Pesanan 2 -> Waktu Proses
// Program Pesanan 1 dan 2 itu berjalan secara bersamaan tanpa menunggu satu sama
// lain
