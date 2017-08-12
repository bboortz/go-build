package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

// var testdir string = os.Getenv("GOPATH") + "/src/github.com/bboortz/go-utils/testdata"
var testdir string = "/app/testdata"
var emptyFile string = testdir + "/empty.txt"
var oneLineFile string = testdir + "/oneline.txt"
var threeLineFile string = testdir + "/threelines.txt"
var unknownFile string = testdir + "/UNKNOWNFILE.txt"

func TestReadLinesEmpty(t *testing.T) {
	assert := assert.New(t)

	result, err := ReadLines(emptyFile)
	assert.Nil(err)
	assert.Nil(result)
	assert.Equal(0, len(result))
}

func TestReadLinesOneLine(t *testing.T) {
	assert := assert.New(t)

	result, err := ReadLines(oneLineFile)
	assert.Nil(err)
	assert.NotNil(result)
	assert.Equal(1, len(result))
}

func TestReadLinesThreeLine(t *testing.T) {
	assert := assert.New(t)

	result, err := ReadLines(threeLineFile)
	assert.Nil(err)
	assert.NotNil(result)
	assert.Equal(3, len(result))
}

func TestReadLinesUnknownFile(t *testing.T) {
	assert := assert.New(t)

	result, err := ReadLines(unknownFile)
	assert.NotNil(err)
	assert.Nil(result)
}
