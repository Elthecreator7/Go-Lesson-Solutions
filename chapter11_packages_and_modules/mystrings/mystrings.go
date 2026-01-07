package mystrings

func Reverse(s string) string {
	letters := []rune(s)
	for i := 0; i < len(letters)/2; i++ {
		j := len(letters) - 1 - i
		letters[i], letters[j] = letters[j], letters[i]
	}
	reversedString := string(letters)
	return reversedString
}
