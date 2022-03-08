package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Cetak "Input: " ke terminal.
	fmt.Fprint(os.Stdout, "Input: ")

	// Buat scanner untuk mendapatkan nilai inputan dari user.
	scanner := bufio.NewScanner(os.Stdin)

	// Scan sampai akhir inputan user.
	scanner.Scan()

	// Ubah nilai input user dari tipe string menjadi int64 dengan basis 10.
	max, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Buat array string kosong.
	result := make([]string, 0)

	// Iterasi dari 1 sampai nilai input user.
	for currentNum := 1; currentNum <= int(max); currentNum++ {
		if val, ok := couldBeDivisible(currentNum); ok {
			result = append(result, val)
		} else {
			// Jika nilai currentNum tidak dapat dibagi dengan 3 dan 5
			// maka tambahkan nilai currentNum ke array result.
			result = append(result, fmt.Sprint(currentNum))
		}
	}

	// Cetak result ke terminal
	fmt.Fprintf(os.Stdout, "Output: %v\n", result)
}

func couldBeDivisible(currentNum int) (string, bool) {
	var divisible bool
	var val string

	switch {
	case currentNum%3 == 0 && currentNum%5 == 0:
		// Jika nilai currentNum dapat dibagi dengan 3 dan 5
		// maka masukkan nilai "Kulina x Food"
		divisible = true
		val = "Kulina x Food"
	case currentNum%3 == 0:
		// Jika nilai currentNum hanya dapat dibagi dengan 3
		// maka masukkan nilai "Kulina"
		divisible = true
		val = "Kulina"
	case currentNum%5 == 0:
		// Jika nilai currentNum hanya dapat dibagi dengan 5
		// maka masukkan nilai "Food"
		divisible = true
		val = "Food"
	}
	return val, divisible
}
