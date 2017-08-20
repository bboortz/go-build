package command

import (
	"os/exec"
)

// Builder is the interface to build a new logger
type Builder interface {
	SuppressStdout() Builder
	SuppressStderr() Builder
	EnableCheckError() Builder
	Build() Command
}

type builder struct {
	command    string
	stdoutOn   bool
	stderrOn   bool
	checkError bool
}

// NewCommand is the create function for the Builder
func NewCommand(commandParam string) Builder {
	return &builder{command: commandParam, stdoutOn: true, stderrOn: true, checkError: false}
}

func (b *builder) SuppressStdout() Builder {
	b.stdoutOn = false
	return b
}

func (b *builder) SuppressStderr() Builder {
	b.stderrOn = false
	return b
}

func (b *builder) EnableCheckError() Builder {
	b.checkError = true
	return b
}

func (b *builder) Build() Command {
	return &command{
		cmd:        exec.Command("/bin/sh", "-c", b.command),
		stdoutOn:   b.stdoutOn,
		stderrOn:   b.stderrOn,
		checkError: b.checkError,
	}
}
