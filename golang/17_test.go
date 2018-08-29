package leetcode

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	digits := "23"
	result := letterCombinations(digits)
	expect := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	if reflect.DeepEqual(digits, result) {
		t.Logf("%s letter combinations is %s", digits, expect)
	} else {
		t.Fatalf("Failed")
	}
}
