package findSubstring

// You are given a string, s, and a list of words, words,
// that are all of the same length. Find all starting indices
// of substring(s) in s that is a concatenation of each word in
// words exactly once and without any intervening

func findSubstring(s string, words []string) []int {
	res := []int{}
	ls, lw := len(s), len(words)
	if lw == 0 || ls == 0 {
		return res
	}

	mp := make(map[string]int)
	for _, v := range words {
		mp[v]++
	}

	wLen := len(words[0])
	window := wLen * lw
	for i := 0; i < wLen; i++ {
		for j := i; j+window <= ls; j = j + wLen {
			tmp := s[j : j+window]
			tp := make(map[string]int)
			for k := lw - 1; k >= 0; k-- {
				// get the word from tmp
				word := tmp[k*wLen : (k+1)*wLen]
				count := tp[word] + 1
				if count > mp[word] {
					j = j + k*wLen
					break
				} else if k == 0 {
					res = append(res, j)
				} else {
					tp[word] = count
				}
			}
		}
	}
	return res
}
