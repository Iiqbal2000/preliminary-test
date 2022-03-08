package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cetak "Input: " ke terminal.
	fmt.Fprint(os.Stdout, "Input: ")

	// Buat scanner untuk mendapatkan nilai inputan dari user.
	scanner := bufio.NewScanner(os.Stdin)

	// Pisahkan masing-masing inputan.
	scanner.Split(bufio.ScanLines)

	fmt.Fprint(os.Stdout, "s = ")
	scanner.Scan()
	// s berisi nilai inputan pertama user
	s := scanner.Text()

	fmt.Fprint(os.Stdout, "t = ")
	scanner.Scan()
	// t berisi nilai inputan kedua user
	t := scanner.Text()

	result := GetDiffChar(s, t)

	fmt.Fprintf(os.Stdout, "Output: %s\n", result)
}

func GetDiffChar(s, t string) string {
	// sBucket menyimpan karakter unik dan jumlah karakter
	// yang muncul pada parameter s.
	sBucket := make(map[rune]int)

	// Iterasi nilai dari parameter s
	for _, char := range s {
		// Jika char sudah ada di dalam sBucket inkremen
		// nilainya.
		if val, ok := sBucket[char]; ok {
			sBucket[char] = val + 1
		} else {
			// Jika char belum ada di sBucket maka masukkan ke-
			// dalamnya.
			sBucket[char] = 1
		}
	}

	var result byte

	// Iterasi nilai dari parameter t
	for _, char := range t {
		// Jika char ada di dalam bucket dan nilainya 0 atau
		// char tidak ada di dalam bucket, masukkan nilai char
		// ke result.
		if val, ok := sBucket[char]; val == 0 || !ok {
			result = byte(char)
			break
		} else {
			// Kurangi nilai char
			sBucket[char]--
		}
	}

	return string(result)
}
