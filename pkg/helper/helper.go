package helper

// Reverse returns a reversed []string
func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// Sum returns the sum of []int
func Sum(array []int) int {
	result := 0

	for _, v := range array {
		result += v
	}
	return result
}
