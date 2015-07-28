package removeDuplicate

// RemoveDuplicate remove duplicate items from slice by setting it to slice2
func RemoveDuplicate(s []string) []string {
	s2 := s[:1]
Loop:
	for i := 1; i < len(s); {
		for j := 0; j < len(s2); {
			if s[i] != s[j] {
				j++
			} else {
				i++
				continue Loop
			}
		}
		s2 = append(s2, s[i])
		i++
	}
	return s2
}
