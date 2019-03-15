package wordBreak

import "strings"

/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
*/

func wordBreak(s string, wordDict []string) bool {
	// transfer list to dict, so that the search time is O(1)
	// dict := make(map[string]bool)
	// for _, word := range wordDict {
	// 	dict[word] = true
	// }
	l := len(s)
	dp := make([]bool, l)

	for i := 0; i < l; i++ {
		for _, word := range wordDict {
			if (i == 0 || dp[i-1]) && strings.HasPrefix(s[i:], word) {
				dp[i+len(word)-1] = true
			}
		}
	}
	return dp[l-1]
}

func wordBreak2(s string, wordDict []string) bool {
	l := len(s)
	d := make([]bool, l)

	for i := 0; i < len(s); i++ {
		for _, w := range wordDict {
			//if at i == 0 consider that we can start break here
			//or if d[i-1] is true meaning that we previously found such word combination that ends on previous position
			// in case if d[i-1] is false - there's no reason to continue searching for a word in dict as there was no combination of words found previously that ends up at [i-1]
			//if the substring starts with word from dictionary -
			if (i == 0 || d[i-1]) && strings.HasPrefix(s[i:], (w)) {
				// mark d at position of the end of a word from a dictionary
				d[i+len(w)-1] = true
			}
		}
	}
	return d[l-1]
}
