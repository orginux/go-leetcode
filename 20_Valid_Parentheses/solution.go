package validparentheses

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	for _, ch := range s {
		if _, ok := brackets[ch]; ok {
			stack = append(stack, ch)
		} else if len(stack) == 0 || brackets[stack[len(stack)-1]] != ch {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
