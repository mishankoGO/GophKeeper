// Package util contains StringInSlice function to check if input string is in slice.
package util

// StringInSlice function returns the presence or absence of a value in the list.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
