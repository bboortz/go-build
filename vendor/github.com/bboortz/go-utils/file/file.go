package file

import (
	"bufio"
	"github.com/bboortz/go-utils/logger"
	"os"
)

var log = logger.NewLogger().Build()

// ReadLines reads a whole file into memory and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}

	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// IsFileExists checks if the file exists
func IsFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
