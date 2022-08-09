package polindrome

import (
	"strings"
)

func Polindrome(kalimat string) bool {

	convertLowerCase := strings.ToLower(kalimat)

	for i, _ := range convertLowerCase {
		lastIndex := (len(convertLowerCase) - 1) - i

		if convertLowerCase[i] != convertLowerCase[lastIndex] {
			return false
		}
	}
	return true
}

func KonsonanOrVocal(huruf string) string {

	convertLowerCase := strings.ToLower(huruf)

	var result string

	if len(convertLowerCase) > 1 {
		result = "Hanya bisa mtestingskan 1 huruf bukan kata atau kalimat"
		return result
	}

	switch convertLowerCase {
	case "a":
		result = "Vocal"
	case "i":
		result = "Vocal"
	case "u":
		result = "Vocal"
	case "e":
		result = "Vocal"
	case "o":
		result = "Vocal"
	default:
		result = "Konsonan"
	}

	return result
}
