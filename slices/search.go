package slices

func IndexOfString(sliceToSearch []string, value string) int {
	for i, val := range sliceToSearch {
		if value == val {
			return i
		}
	}
	return -1
}
