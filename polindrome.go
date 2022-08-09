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
