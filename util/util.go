package util

func findIndex(array [4]string, value string) int {
	for p, v := range array {
		if v == value {
			return p
		}
	}
	return -1
}
