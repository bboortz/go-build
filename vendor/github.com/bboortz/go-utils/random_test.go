package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	//	"math"
	"math/rand"
	"testing"
)

func init() {
	Seed()
}

func TestGenSeed(t *testing.T) {
	a := assert.New(t)

	result := GenSeed()
	a.NotNil(result)
}

func TestRandomString(t *testing.T) {
	a := assert.New(t)

	result := RandomString(10)
	a.NotNil(result)
	a.Equal(10, len(result))
}

func TestRandomInt(t *testing.T) {
	a := assert.New(t)

	result := RandomInt(10, 100)
	a.NotNil(result)
	a.Equal(true, result >= 10)
	a.Equal(true, result <= 100)
}

func TestRandomIntWithSeed1(t *testing.T) {
	a := assert.New(t)
	rand.Seed(int64(1))

	result := RandomInt(10, 100)
	a.NotNil(result)
	a.Equal(true, result >= 10)
	a.Equal(true, result <= 100)
	a.Equal(51, result)
}
