package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readTemperature() (float32, error) {
	reader := bufio.NewReader(os.Stdin)
	suhuInput, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	trimInput := strings.TrimSpace(suhuInput)

	suhuCelcius, err := strconv.ParseFloat(trimInput, 32)

	if err != nil {
		return 0, errors.New("Input tidak valid, hanya menerima angka!")
	}

	return float32(suhuCelcius), nil
}

func main() {
	fmt.Print("Masukkan suhu dalam Celcius: ")
	suhuCelcius, error := readTemperature()

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Suhu dalam Celcius: ", suhuCelcius)
	}
}
