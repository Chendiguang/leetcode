package valid_Palindrome

import (
	"unicode"
)

/*
需要考虑大小写，需要考虑数字，不考虑标点符号
*/
func isPalindrome(s string) bool {
	var letters []rune
	for _, c := range s {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			letters = append(letters, unicode.ToLower(c))
		}
	}
	for i, j := 0, len(letters)-1; i <= j; i, j = i+1, j-1 {
		if letters[i] != letters[j] {
			return false
		}
	}
	return true
}
