package lightEnigma

// EnigmaON - shift k bytes foward
func EnigmaON(word string, k uint8) string {
	var enigON []byte
	for letter := range word {
		enigON = append(enigON, (word[letter] + k))
	}
	enigONstr := string(enigON)
	return enigONstr
}

// EnigmaOFF - shift k bytes backward
func EnigmaOFF(enigON string, k uint8) string {
	var enigOFF []byte
	for letter := range enigON {
		enigOFF = append(enigOFF, enigON[letter]-k)
	}
	enigOFFstr := string(enigOFF)
	return enigOFFstr
}

//
