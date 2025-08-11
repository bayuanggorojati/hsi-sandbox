package main

import (
	"errors"
	"fmt"
)

type rawData struct {
	data         []int
	channel      chan int
	channelError chan error
}

func hitungData(rd rawData) {
	if len(rd.data) == 0 {
		rd.channel <- 0 // Kirim 0 jika data kosong
		rd.channelError <- errors.New("Deret nya tidak boleh kosong")
		return
	} // Kirim 0 jika data kosong
	total := 0
	for _, value := range rd.data {
		total += value
	}
	rd.channel <- total // Kirim total ke channel
}

func main() {
	angkaGenap := []int{2, 4, 6, 8, 10}
	angkaGanjil := []int{1, 3, 5, 7, 9}

	c1 := make(chan int)
	c2 := make(chan int)

	// c1Error := make(chan error)
	// c2Error := make(chan error)

	deret1 := rawData{data: angkaGenap, channel: c1}  // 5 ms
	deret2 := rawData{data: angkaGanjil, channel: c2} // 12 ms

	go hitungData(deret1)
	go hitungData(deret2)

	totalAngkaGenap := <-c1  // Tunggu dan ambil total dari channel c1
	totalAngkaGanjil := <-c2 // Tunggu dan ambil total dari channel c2

	fmt.Println("Total Angka Genap:", totalAngkaGenap)
	fmt.Println("Total Angka Ganjil:", totalAngkaGanjil)

	fmt.Println("Total Keseluruhan:", totalAngkaGenap+totalAngkaGanjil) // Total keseluruhan dari kedua channel
}
