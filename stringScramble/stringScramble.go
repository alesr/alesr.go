// Using the !JavaScript (but Go =) language, have the function StringScramble(str1,str2)
// take both parameters being passed and return the string true if a portion
// of str1 characters can be rearranged to match str2, otherwise return the
// string false. For example: if str1 is "rkqodlw" and str2 is "world" the
// output should return true. Punctuation and symbols will not be entered
// with the parameters.

package main

import "fmt"

func main() {
	str1, str2 := "hello", "ello"
	StringScramble(str1, str2)
}

// StringScramble checks if str2 is on str1
func StringScramble(str1, str2 string) bool {
	str1Len, str2Len := len(str1), len(str2)

	// A word may not contain another one,
	// if the first is shorter than the second =P
	if str1Len < str2Len {
		fmt.Printf("FALSE\n")
		return false
	}
	count := 0

	// For each letter on str1, lets compare with the all letter
	// on str2.
	for i := 0; i < len(str1); {
		for j := 0; j < len(str2); {
			if str1[i] == str2[j] {
				count++
				i++
			}
			j++
		}
		i++
	}

	//We could make the loop more complicated, but we can guarantee that
	//if we find the len(str2) of correspondence letters on str1.
	// Then str1 must contain the str2.
	if count == str2Len {
		fmt.Printf("TRUE\n")
		return true
	}

	return false
}
