package main_test

import (
	"fmt"
	"testing"

	roman "github.com/Iiqbal2000/roman"
)

func TestConvertToRoman(t *testing.T) {
	cases := []struct{
		roman string
		want int
	} {
		{roman: "XII", want: 12},
		{roman: "LIV", want: 54},
		{roman: "MMXXII", want: 2022},
		{roman: "I", want: 1},
		{roman: "II", want: 2},
		{roman: "III", want: 3},
		{roman: "IV", want: 4},
		{roman: "V", want: 5},
		{roman: "VI", want: 6},
		{roman: "VII", want: 7},
		{roman: "VIII", want: 8},
		{roman: "IX", want: 9},
		{roman: "X", want: 10},
		{roman: "XIV", want: 14},
		{roman: "XVIII", want: 18},
		{roman: "XX", want: 20},
		{roman: "XXXIX", want: 39},
		{roman: "XL", want: 40},
		{roman: "XLVII", want: 47},
		{roman: "XLIX", want: 49},
		{roman: "L", want: 50},
		{roman: "C", want: 100},
		{roman: "XC", want: 90},
		{roman: "CD", want: 400},
		{roman: "D", want: 500},
		{roman: "CM", want: 900},
		{roman: "M", want: 1000},
		{roman: "MCMLXXXIV", want: 1984},
		{roman: "MMMCMXCIX", want: 3999},
		{roman: "MMXIV", want: 2014},
		{roman: "MVI", want: 1006},
		{roman: "DCCXCVIII", want: 798},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test-%d (%s)", i+1, c.roman), func(t *testing.T) {
			got := roman.ConvertToArabic(c.roman)
			if got != c.want {
				t.Errorf("Got %d, want %d", got, c.want)
			}
		})
	}
}