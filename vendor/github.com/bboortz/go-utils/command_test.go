package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestExecCommandTrue(t *testing.T) {
	assert := assert.New(t)

	exitCode, stdoutStr, stderrStr := ExecCommand("/bin/true")
	assert.Equal(0, exitCode)
	assert.Equal("", stdoutStr)
	assert.Equal("", stderrStr)
}

func TestExecCommandFalse(t *testing.T) {
	assert := assert.New(t)

	exitCode, stdoutStr, stderrStr := ExecCommandWithoutErrCheck("/bin/false")
	assert.Equal(1, exitCode)
	assert.Equal("", stdoutStr)
	assert.Equal("", stderrStr)
}

func TestExecCommandEcho(t *testing.T) {
	assert := assert.New(t)

	exitCode, stdoutStr, stderrStr := ExecCommand("/bin/echo test")
	assert.Equal(0, exitCode)
	assert.Equal("test\n", stdoutStr)
	assert.Equal("", stderrStr)
}
