package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestExecCommandTrue(t *testing.T) {
	a := assert.New(t)

	exitCode, stdoutStr, stderrStr := ExecCommand("/bin/true")
	a.Equal(0, exitCode)
	a.Equal("", stdoutStr)
	a.Equal("", stderrStr)
}

func TestExecCommandFalse(t *testing.T) {
	a := assert.New(t)

	exitCode, stdoutStr, stderrStr := ExecCommandWithoutErrCheck("/bin/false")
	a.Equal(1, exitCode)
	a.Equal("", stdoutStr)
	a.Equal("", stderrStr)
}

func TestExecCommandEcho(t *testing.T) {
	a := assert.New(t)

	exitCode, stdoutStr, stderrStr := ExecCommand("/bin/echo test")
	a.Equal(0, exitCode)
	a.Equal("test\n", stdoutStr)
	a.Equal("", stderrStr)
}

func TestExecCommandWithOutput(t *testing.T) {
	a := assert.New(t)

	exitCode, err := ExecCommandWithOutput("/bin/echo test", true, true)
	a.Equal(0, exitCode)
	a.Nil(err)
}

func TestExecCommandWithOutputBad(t *testing.T) {
	a := assert.New(t)

	exitCode, err := ExecCommandWithOutput("/bin/cat doesnotexist", true, false)
	a.Equal(1, exitCode)
	a.NotNil(err)
}
