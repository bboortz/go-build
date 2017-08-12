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
	assert := assert.New(t)

	var result int64 = GenSeed()
	assert.NotNil(result)
}

func TestRandomString(t *testing.T) {
	assert := assert.New(t)

	var result string = RandomString(10)
	assert.NotNil(result)
	assert.Equal(10, len(result))
}

func TestRandomInt(t *testing.T) {
	assert := assert.New(t)

	var result int = RandomInt(10, 100)
	assert.NotNil(result)
	assert.Equal(true, result >= 10)
	assert.Equal(true, result <= 100)
}

func TestRandomIntWithSeed1(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(int64(1))

	var result int = RandomInt(10, 100)
	assert.NotNil(result)
	assert.Equal(true, result >= 10)
	assert.Equal(true, result <= 100)
	assert.Equal(51, result)
}
