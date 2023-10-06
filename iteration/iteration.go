package iteration

func Repeat(text string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += text
	}
	return repeated
}
