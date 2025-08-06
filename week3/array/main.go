package main

import "fmt"

//Array struktur data yang mempunyai ukuran tetap
//Tipe data array itu selalu sama

func main() {
	//Cara membuat array metode 1
	var arrayAngka [10]int // index selalu dimulai dari 0 sampai dengan 9
	//Cara membuat array metode 2
	arrayAngka1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Data dari arrayAngka1 index 2:", arrayAngka1[2]) //mencetak angka 3

	// Cara mengisi data array dari metode 1
	arrayAngka[0] = 2
	arrayAngka[1] = 4
	arrayAngka[2] = 6
	arrayAngka[3] = 8
	arrayAngka[4] = 10
	arrayAngka[5] = 12
	fmt.Println("Data dari arrayAngka index 4:", arrayAngka[4]) // mencetaka angka 10

	// Cara membuat array metode 3
	arrayAngka2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}      // panjang array otomatis ditentukan
	fmt.Println("Panjang Array arrayAngka2:", len(arrayAngka2)) // mencetak panjang array 10
}
