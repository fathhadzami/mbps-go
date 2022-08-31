package main

import (
	"fmt"
)

func isContains(listAlphabet []string, alphabet string) bool {
	for _, v := range listAlphabet {
		if v == alphabet {
			return true
		}
	}
	return false
}

func isConsonant(input string) bool {
	Consonant := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z"}
	return isContains(Consonant, input)
}

func isVocal(input string) bool {
	Vocal := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	return isContains(Vocal, input)
}

func isNumber(input string) bool {
	Number := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return isContains(Number, input)
}
func checkConsonantOrVocal(input string) string {
	if isConsonant(input) {
		return "Consonant"
	} else if isVocal(input) {
		return "Vocal"
	} else if isNumber(input) {
		return "Hanya diperbolehkan Huruf"
	}
	return "Hanya diperbolehkan 1 Huruf"
}

func main() {
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(checkConsonantOrVocal(input))
}
