package maputil

import ()

// IndexOf retrieves the index of a string inside a map[int]string
func IndexOf(word string, data map[int]string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
