package reverseString

// ReverseString - output the reverse string of a given string s
func ReverseString(s string) string {
	revStr := make([]byte, len(s))
	for i := range s {
		revStr[i] += s[len(s)-1-i]
	}
	return string(revStr)
}
