package iteration

const repeatCount = 5

func Repeat(char string, count int) string {
	var repeated string
	if count == 0 {
		for i := 0; i < repeatCount; i++ {
			repeated += char
		}
		return repeated
	}
	repeatCount := count
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
