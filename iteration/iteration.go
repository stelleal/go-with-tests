package iteration

import "strings"

// Repeat takes a character and returns a string with it repeated 5 times
func Repeat(character string, count int) string {
	// var repeated string
	// for i := 0; i < count; i++ {
	// 	repeated += character
	// }
	return strings.Repeat(character, count)
}
