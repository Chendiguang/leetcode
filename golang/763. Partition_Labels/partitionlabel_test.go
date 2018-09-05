package Partition_Labels

import (
	"fmt"
	"testing"
)

func ExamplepartitionLabels() {
	S := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(S))
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_partitionLabels(t *testing.T) {
	var tests = []struct {
		input string
		want  []int
	}{
		{
			"ababcbacadefegdehijhklij",
			[]int{9, 7, 8},
		},
		{
			"abcdefg",
			[]int{1, 1, 1, 1, 1, 1, 1},
		},
		// ...
	}

	for _, test := range tests {
		if got := partitionLabels(test.input); !intSliceEqual(got, test.want) {
			t.Errorf("partitionLabels(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
