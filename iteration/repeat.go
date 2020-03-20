package iteration

const repeatedCount = 5

// Repeat repeat a give character 5 times.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}
	return repeated
}
