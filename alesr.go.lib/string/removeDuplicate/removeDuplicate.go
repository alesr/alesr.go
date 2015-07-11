package removeDuplicate

// RemoveDuplicate remove duplicate items from array setting it to arr2
func RemoveDuplicate(arr []string) []string {
	arr2 := arr[:1]
Loop:
	for i := 1; i < len(arr); {
		for j := 0; j < len(arr2); {
			if arr[i] != arr[j] {
				j++
			} else {
				i++
				continue Loop
			}
		}
		arr2 = append(arr2, arr[i])
		i++
	}
	return arr2
}
