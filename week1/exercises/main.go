package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celciusToFahrenheit(temp float32) float32 {
	return temp*9/5 + 32
}

func celciusToReamur(temp float32) float32 {
	return temp * 4 / 5
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan suhu dalam Celcius: ")
	suhuInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error dalam membaca input suhu: ", err)
		return
	}

	trimInput := strings.TrimSpace(suhuInput)

	suhuCelcius, err := strconv.ParseFloat(trimInput, 32)

	if err != nil {
		fmt.Println("Input tidak valid, hanya menerima angka")
		return
	}

	//var suhu float32 = 23

	suhuReamur := celciusToReamur(float32(suhuCelcius))
	suhuFahrenheit := celciusToFahrenheit(float32(suhuCelcius))

	fmt.Println("Suhu dalam Reamur: ", suhuReamur)
	fmt.Println("Suhu dalam Fahrenheit: ", suhuFahrenheit)
}
