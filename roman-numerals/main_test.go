package roman_numerals

import (
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvertToRoman(tc.arabic)

			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

var tests = []struct {
	name   string
	arabic uint16
	want   string
}{
	{
		name:   "1 equals I",
		arabic: 1,
		want:   "I",
	},
	{
		name:   "2 equals II",
		arabic: 2,
		want:   "II",
	},
	{
		name:   "4 equals IV",
		arabic: 4,
		want:   "IV",
	},
	{
		name:   "5 equals V",
		arabic: 5,
		want:   "V",
	},
	{
		name:   "8 equals VIII",
		arabic: 8,
		want:   "VIII",
	},
	{
		name:   "9 equals IX",
		arabic: 9,
		want:   "IX",
	},
	{
		name:   "10 equals X",
		arabic: 10,
		want:   "X",
	},
	{
		name:   "11 equals XI",
		arabic: 11,
		want:   "XI",
	},
	{
		name:   "14 equals XIV",
		arabic: 14,
		want:   "XIV",
	},
	{
		name:   "18 equals XIX",
		arabic: 18,
		want:   "XVIII",
	},
	{
		name:   "20 equals XX",
		arabic: 20,
		want:   "XX",
	},
	{
		name:   "39 equals XXXIX",
		arabic: 39,
		want:   "XXXIX",
	},
	{
		name:   "40 equals XL",
		arabic: 40,
		want:   "XL",
	},
	{
		name:   "47 equals XLVII",
		arabic: 47,
		want:   "XLVII",
	},
	{
		name:   "49 equals XLIX",
		arabic: 49,
		want:   "XLIX",
	},
	{
		name:   "50 equals L",
		arabic: 50,
		want:   "L",
	},
	{
		name:   "54 equals LIV",
		arabic: 54,
		want:   "LIV",
	},
	{
		name:   "100 equals C",
		arabic: 100,
		want:   "C",
	},
	{
		name:   "103 equals CIII",
		arabic: 103,
		want:   "CIII",
	},
	{
		name:   "400 equals CD",
		arabic: 400,
		want:   "CD",
	},
	{
		name:   "500 equals D",
		arabic: 500,
		want:   "D",
	},
	{
		name:   "900 equals CM",
		arabic: 900,
		want:   "CM",
	},
	{
		name:   "1000 equals M",
		arabic: 1000,
		want:   "M",
	},
	{
		name:   "1984 equals MCMLXXXIV",
		arabic: 1984,
		want:   "MCMLXXXIV",
	},
}
