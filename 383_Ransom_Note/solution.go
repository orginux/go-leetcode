package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	if ransomNote == magazine {
		return true
	}
	if len(ransomNote) > len(magazine) {
		return false
	}

	letters := map[rune]int{}
	for _, ch := range magazine {
		letters[ch]++
	}

	for _, ch := range ransomNote {
		if letters[ch] == 0 {
			return false
		}
		letters[ch]--
	}
	return true
}
