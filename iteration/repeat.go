package iteration

// Repeat takes a character and repeats it five times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}
