package sliceutil

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestIndexOfGood(t *testing.T) {
	a := assert.New(t)

	var testdata = []string{
		"CRITICAL",
		"ERROR",
	}

	i := IndexOf("CRITICAL", testdata)
	a.Equal(0, i)
}

func TestIndexOfBad(t *testing.T) {
	a := assert.New(t)

	var testdata = []string{
		"CRITICAL",
		"ERROR",
	}

	i := IndexOf("INFO", testdata)
	a.Equal(-1, i)
}

func TestIndexOfEmpty(t *testing.T) {
	a := assert.New(t)

	var testdata = []string{}

	i := IndexOf("CRITICAL", testdata)
	a.Equal(-1, i)
}
