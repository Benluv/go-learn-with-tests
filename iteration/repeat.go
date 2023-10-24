package iteration

import "strings"

func Repeat(character string, repeat int) string {
	return strings.Repeat(character, repeat)
}

func CountRepeat(character string, count int) (counted int) {
	var repeated string
	if count == 0 {
		count = 8
	}
	repeated = strings.Repeat(character, count)
	counted = strings.Count(repeated, character)
	return
}
