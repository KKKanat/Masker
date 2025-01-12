package utils

import (
	"sync"
)

func RuneLength(args string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

func Find(text string, ch chan<- string, wg *sync.WaitGroup) {
	pattern := "http://"
	patternLength := RuneLength(pattern)
	needToMask := false
	textInRunes := []rune(text)
	defer wg.Done()
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
	ch <- string(textInRunes)
}
