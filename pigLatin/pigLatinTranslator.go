package pigLatin

import "strings"

// "egg" -> "eggway"
// "pig" -> "igay"

func PigTranslator(s string) string {
	word := strings.ToLower(s)
	var latin string
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u':
		latin = word + "way"
	default:
		latin = word[1:] + string(word[0]) + "ay"
	}
	return latin
}
