package stringutil

import (
	"github.com/bboortz/go-utils/logger"
	"strconv"
)

var log = logger.NewLogger().Build()

// CheckEmpty checks if a variable is empty. If the string is empty the key will be logged out and exit imediately.
func CheckEmpty(key string, value string) {
	if value == "" {
		log.Fatal("variable with key <" + key + "> is empty.")
	}
}

// ConvertStringArrayToIntArray converts a String Array to and Int Array
func ConvertStringArrayToIntArray(arr []string) ([]int, error) {
	var result = []int{}
	var err error

	for _, i := range arr {
		var j int
		j, err = strconv.Atoi(i)
		if err != nil {
			return result, err
		}
		result = append(result, j)
	}

	return result, err

}
