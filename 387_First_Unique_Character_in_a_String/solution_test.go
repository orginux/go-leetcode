package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	characters := make(map[rune]int)

	for _, ch := range s {
		characters[ch]++
	}

	for i, ch := range s {
		if characters[ch] == 1 {
			return i
		}
	}

	return -1
}
