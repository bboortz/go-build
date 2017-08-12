package maps

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestIndexOfGood(t *testing.T) {
	assert := assert.New(t)

	var testdata = map[int]string{
		0: "string1",
		1: "string2",
		2: "string3",
	}

	i := IndexOf("string1", testdata)
	assert.Equal(0, i)
}

func TestIndexOfBad(t *testing.T) {
	assert := assert.New(t)

	var testdata = map[int]string{
		0: "string1",
		1: "string2",
		2: "string3",
	}

	i := IndexOf("string4", testdata)
	assert.Equal(-1, i)
}

func TestIndexOfEmpty(t *testing.T) {
	assert := assert.New(t)

	var testdata = map[int]string{}

	i := IndexOf("string1", testdata)
	assert.Equal(-1, i)
}
