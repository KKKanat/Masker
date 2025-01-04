package utils

func RuneLength(args string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

func Find(text string) string {
	pattern := "http://"
	patternLength := RuneLength(pattern)
	needToMask := false
	textInRunes := []rune(text)
	for i, r := range textInRunes {
		if r == ' ' {
			needToMask = false
		}
		if needToMask {
			textInRunes[i] = '*'
		}
		if !needToMask && i+1 >= patternLength {
			if pattern == text[i+1-patternLength:i+1] {
				needToMask = true
			}
		}
	}
	return string(textInRunes)
}