package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StructTemp struct {
	tempCelcius float32
	location    string
}

func celciusToFahrenheit(temp *StructTemp) float32 {
	return temp.tempCelcius*9/5 + 32
}

func celciusToReamur(temp *StructTemp) float32 {
	return temp.tempCelcius * 4 / 5
}

func tempClass(suhuCelcius float32) string {
	if suhuCelcius < 18 {
		return "dingin"
	} else if suhuCelcius > 25 {
		return "panas"
	}
	return "hangat"
}

func readUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(userInput), nil
}

func isString(location string) bool {
	_, err := strconv.Atoi(location)
	if err != nil {
		return true
	}
	return false
}

func stringToFloat(tempInput string) (float32, error) {
	suhuCelcius, err := strconv.ParseFloat(tempInput, 32)

	if err != nil {
		return 0, errors.New("Input tidak valid, hanya menerima angka!")
	}

	return float32(suhuCelcius), nil
}

func main() {
	var tempInput StructTemp

	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan lokasi pengukuran suhu : ")

	lokasi, err := readUserInput()
	if err != nil {
		fmt.Println("Error membaca lokasi pengukuran suhu: ", err)
	}

	if !isString(lokasi) {
		fmt.Println("Input lokasi tidak valid, tidak menerima angka/numeric")
		return
	}

	tempInput.location = lokasi

	fmt.Print("Masukkan suhu dalam Celcius: ")

	tempInputString, err := readUserInput()
	if err != nil {
		fmt.Println("Error membaca input suhu: ", err)
	}

	suhuCelcius, err := stringToFloat(tempInputString)
	//var suhu float32 = 23

	if err != nil {
		fmt.Println(err)
	} else {
		tempInput.tempCelcius = suhuCelcius
		fmt.Println("Suhu di", tempInput.location, "adalah", tempClass(suhuCelcius))
		suhuReamur := celciusToReamur(&tempInput)
		suhuFahrenheit := celciusToFahrenheit(&tempInput)

		fmt.Println("Suhu di", tempInput.location, "dalam Reamur: ", suhuReamur)
		fmt.Println("Suhu di", tempInput.location, "dalam Fahrenheit: ", suhuFahrenheit)
	}
}
