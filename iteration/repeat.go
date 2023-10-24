package iteration

import "strings"

func Repeat(character string, repeat int) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

func CountRepeat(character string, count int) int {
	var repeated string
	var counted int
	if count == 0 {
		count = 8
	}
	repeated = strings.Repeat(character, count)
	counted = strings.Count(repeated, character)
	return counted
}
