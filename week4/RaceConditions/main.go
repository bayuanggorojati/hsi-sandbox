package main

import "fmt"

var hitung int // sumber konflik data (data race situation)

func tambah() {
	hitung++
}

func main() {
	for i := 0; i < 100; i++ {
		go tambah()
		// 1 : 2 goroutine berjalan
		// 2 : 1 goroutine berjalan
		// 3 : 2 goroutine berjalan
	}

	fmt.Println("Hasil akhir: ", hitung)
}
