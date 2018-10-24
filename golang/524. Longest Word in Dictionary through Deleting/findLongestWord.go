package findLongestWord

import (
	"strings"
)

func findLongestWord(s string, d []string) string {
	max_str := ""
	for _, str := range d {
		if isSubsequence(str, s) {
			lm, ls := len(max_str), len(str)
			if lm < ls || (lm == ls && strings.Compare(max_str, str) > 0) {
				max_str = str
			}
		}
	}
	return max_str
}

func isSubsequence(x, y string) bool {
	j := 0
	for i := 0; i < len(y) && j < len(x); i++ {
		if x[j] == y[i] {
			j++
		}
	}
	return j == len(x)
}
