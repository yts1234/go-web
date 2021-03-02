package iteration

const repeatCount = 5

// Repeat will return repeated n count word
func Repeat(char string, count int) string {
	var repeated string
	if count == 0 {
		for i := 0; i < repeatCount; i++ {
			repeated += char
		}
		return repeated
	}

	for i := 0; i < count; i++ {
		repeated += char
	}
	return repeated
}
