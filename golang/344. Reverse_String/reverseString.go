package Reverse_String

/*
Example 1:
Input: "hello"
Output: "olleh"

Example 2:
Input: "A man, a plan, a canal: Panama"
Output: "amanaP :lanac a ,nalp a ,nam A"
*/

// Time Complexity: O(n), Space Complexity: O(n)
func reverseString(s string) string {
	l := len(s)
	ans := make([]byte, l)
	for i := 0; i < l; i++ {
		ans[l-i-1] = s[i]
	}
	return string(ans)
}

// Time Complexity: O(n/2), Space Complexity: O(n)
// double pointer
func reverseString2(s string) string {
	runes := []rune(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}
