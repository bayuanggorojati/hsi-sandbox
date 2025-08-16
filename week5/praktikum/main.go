// main.go
// Entry point dari aplikasi kita
// Special Package dari bahasa go
package main

import (
	"fmt"
	"praktikum/calculator"
	"praktikum/stringgenerator"
)

func main() {
	// Memanggil function TambahAngka dari package calculator
	hasil := calculator.TambahAngka(2, 4)
	fmt.Println("Hasil Penjumlahan:", hasil)

	// Memanggil function KurangAngka dari package calculator
	hasilKurang := calculator.KurangAngka(10, 5)
	fmt.Println("Hasil Pengurangan:", hasilKurang)

	pesanBaru := stringgenerator.MencetakSalam("Darizki")
	fmt.Println(pesanBaru)
}
