package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
}

var testdir = os.Getenv("GOPATH") + "/src/github.com/bboortz/go-utils/testdata"

//var testdir string = "/app/testdata"
var emptyFile = testdir + "/empty.txt"
var oneLineFile = testdir + "/oneline.txt"
var threeLineFile = testdir + "/threelines.txt"
var unknownFile = testdir + "/UNKNOWNFILE.txt"

func TestReadLinesEmpty(t *testing.T) {
	a := assert.New(t)

	result, err := ReadLines(emptyFile)
	a.Nil(err)
	a.Nil(result)
	a.Equal(0, len(result))
}

func TestReadLinesOneLine(t *testing.T) {
	a := assert.New(t)

	result, err := ReadLines(oneLineFile)
	a.Nil(err)
	a.NotNil(result)
	a.Equal(1, len(result))
}

func TestReadLinesThreeLine(t *testing.T) {
	a := assert.New(t)

	result, err := ReadLines(threeLineFile)
	a.Nil(err)
	a.NotNil(result)
	a.Equal(3, len(result))
}

func TestReadLinesUnknownFile(t *testing.T) {
	a := assert.New(t)

	result, err := ReadLines(unknownFile)
	a.NotNil(err)
	a.Nil(result)
}
