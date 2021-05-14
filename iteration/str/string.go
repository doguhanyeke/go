package str

import "strings"

func CompareStrings(a string, b string) int {
	result := strings.Compare(a, b)
	return result
}
