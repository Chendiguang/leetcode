package Reverse_Vowels_of_a_String

// 8ms
func reverseVowels(s string) string {
	runes := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		for !isVowel(runes[r]) && l < r {
			r--
		}

		for !isVowel(runes[l]) && l < r {
			l++
		}
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}

func isVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

//+! solution 2

// byte is faster than rune, 4ms
func reverseVowels2(s string) string {
	b := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for !isVowel2(b[r]) && l < r {
			r--
		}

		for !isVowel2(b[l]) && l < r {
			l++
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isVowel2(r byte) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}
