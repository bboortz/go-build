package stack

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStack(t *testing.T) {
	a := assert.New(t)

	result := GetStack()
	a.NotNil(result)
}

func TestPrintStack(t *testing.T) {
	PrintStack()
}

func TestTrace(t *testing.T) {
	Trace()
}

func TestGetCurrentMethodName(t *testing.T) {
	a := assert.New(t)

	result := GetCurrentMethodName()
	a.NotNil(result)
	a.Equal("TestGetCurrentMethodName", result)
}

func TestGetCallingMethodName(t *testing.T) {
	a := assert.New(t)

	result := GetCallingMethodName()
	a.NotNil(result)
	a.Equal("tRunner", result)
}
