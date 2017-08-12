package maps

import ()

func IndexOf(word string, data map[int]string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
