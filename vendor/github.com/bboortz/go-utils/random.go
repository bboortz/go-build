package utils

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	Seed()
}

// Seed the random generator before using it
func Seed() {
	rand.Seed(GenSeed())
	rand.Seed(RandomInt64(0, math.MaxInt64))
}

// GenSeed generates a random seed
func GenSeed() int64 {
	return time.Now().UTC().UnixNano() + int64(RandomInt(0, 9999999))
}

// RandomString generates a random string
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

// RandomInt generates a random int value
func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// RandomInt64 generates a random int64 value
func RandomInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min)
}
