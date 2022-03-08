package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "Input: s = ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	result := ConvertToArabic(s)

	fmt.Fprintf(os.Stdout, "Output: %d\n", result)
}

// RomanArabic berisi tujuh simbol roman beserta
// nilai arabicnya.
var RomanArabic = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// ConvertToArabic mengubah simbol roman menjadi simbol
// arabic.
func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		romanChar := roman[i]
		arabicSymbol := RomanArabic[romanChar]

		if validSubtractor(i, romanChar, roman) {
			total -= arabicSymbol
		} else {
			total += arabicSymbol
		}
	}

	return total
}

// validSubtractor memeriksa romanChar valid atau tidak
// jika dilakukan operasi pengurangan.
func validSubtractor(index int, romanChar uint8, roman string) bool {
	if i := index+1; i < len(roman) {
		switch {
		case romanChar == 'I' && (roman[i] == 'V' || roman[i] == 'X'):
			return true
		case romanChar == 'X' && (roman[i] == 'L' || roman[i] == 'C'):
			return true
		case romanChar == 'C' && (roman[i] == 'D' || roman[i] == 'M'):
			return true
		}
	}

	return false
}
