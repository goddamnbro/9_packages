package homework

import "fmt"

func CompressWord(s string) interface{} {
	var counter int
	var lastRune rune
	var compressed string

	runes := []rune(s)
	for pos, rune := range runes {

		// first char
		if pos == 0 {
			lastRune = rune
			counter += 1
			continue
		}

		// last char
		if pos == len(runes)-1 {
			// last char equals last char in string
			if lastRune == rune {
				counter += 1
				return fmt.Sprintf("%s%c%d", compressed, rune, counter)
			}

			return fmt.Sprintf("%s%c%d%c%d", compressed, lastRune, counter, rune, 1)
		}

		// prev char equals current char
		if lastRune != rune {
			compressed += fmt.Sprintf("%c%d", lastRune, counter)
			lastRune = rune
			counter = 0
		}

		counter += 1
	}

	return compressed
}
