package iteration

// Repeat takes a character and count, and repeats the character as many times as defined by count
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
