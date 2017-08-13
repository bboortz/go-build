package stringutil

import (
	"github.com/bboortz/go-utils/logger"
)

var log = logger.NewLogger().Build()

// CheckEmpty checks if a variable is empty. If the string is empty the key will be logged out and exit imediately.
func CheckEmpty(key string, value string) {
	if value == "" {
		log.Fatal("variable with key <" + key + "> is empty.")
	}
}
