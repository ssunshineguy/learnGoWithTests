package iteration

func Repeat(str string, repeatCount int) (repited_string string) {
	repited_string = ""
	for i := 0; i < repeatCount; i++ {
		repited_string += str
	}
	return
}
