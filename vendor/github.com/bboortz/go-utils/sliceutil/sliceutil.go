package sliceutil

import ()

// IndexOf retrieves the index of a string inside a slice []string
func IndexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
