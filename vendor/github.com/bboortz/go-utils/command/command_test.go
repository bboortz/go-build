package command

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func init() {
}

func TestBuildCommand(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/true").Build()
	a.NotNil(cmd)
}

func TestBuildCommandAllOptions(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/true").SuppressStdout().SuppressStderr().EnableCheckError().Build()
	a.NotNil(cmd)
}

func TestCommandRun(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /etc/hosts").Build()
	a.NotNil(cmd)
	exitCode, stdout, stderr, err := cmd.Run()
	a.Equal(0, exitCode)
	a.NotEqual("", stdout)
	a.Equal("", stderr)
	a.Nil(err)
}

func TestCommandSuppressedOutputRun(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /etc/hosts").SuppressStdout().SuppressStderr().Build()
	a.NotNil(cmd)
	exitCode, stdout, stderr, err := cmd.Run()
	a.Equal(0, exitCode)
	a.Equal("", stdout)
	a.Equal("", stderr)
	a.Nil(err)
}

func TestCommandRunBad(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /tmp/DOESNOTEXIST").Build()
	a.NotNil(cmd)
	exitCode, stdout, stderr, err := cmd.Run()
	a.Equal(1, exitCode)
	a.Equal("", stdout)
	a.NotEqual("", stderr)
	a.NotNil(err)
}

func TestCommandRunEnableCheckErrorBad(t *testing.T) {
	a := assert.New(t)

	if os.Getenv("TESTRUN") == "1" {
		cmd := NewCommand("/bin/false").EnableCheckError().Build()
		_, _, _, _ = cmd.Run()
		log.Fatal("Code cannot be reached!")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCommandRunEnableCheckErrorBad")
	cmd.Env = append(os.Environ(), "TESTRUN=1")
	err := cmd.Run()
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	a.NotNil(err)
	a.Equal(1, exitCode)
}

func TestCommandRunSuppressOutputBad(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /tmp/DOESNOTEXIST").SuppressStdout().SuppressStderr().Build()
	a.NotNil(cmd)
	exitCode, stdout, stderr, err := cmd.Run()
	a.Equal(1, exitCode)
	a.Equal("", stdout)
	a.Equal("", stderr)
	a.NotNil(err)
}

func TestCommandRunWithLiveOutput(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /etc/hosts").Build()
	a.NotNil(cmd)
	exitCode, err := cmd.RunWithLiveOutput()
	a.Equal(0, exitCode)
	a.Nil(err)
}

func TestCommandSuppressedOutputRunWithLiveOutput(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /etc/hosts").SuppressStdout().SuppressStderr().Build()
	a.NotNil(cmd)
	exitCode, err := cmd.RunWithLiveOutput()
	a.Equal(0, exitCode)
	a.Nil(err)
}

func TestCommandRunWithLiveOutputBad(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /tmp/DOESNOTEXIST").Build()
	a.NotNil(cmd)
	exitCode, err := cmd.RunWithLiveOutput()
	a.Equal(1, exitCode)
	a.NotNil(err)
}

func TestCommandRunWithLiveOutputEnableCheckErrorBad(t *testing.T) {
	a := assert.New(t)

	if os.Getenv("TESTRUN") == "1" {
		cmd := NewCommand("/bin/false").EnableCheckError().Build()
		_, _ = cmd.RunWithLiveOutput()
		log.Fatal("Code cannot be reached!")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCommandRunWithLiveOutputEnableCheckErrorBad")
	cmd.Env = append(os.Environ(), "TESTRUN=1")
	err := cmd.Run()
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	a.NotNil(err)
	a.Equal(1, exitCode)
}

func TestCommandSuppressOutputRunWithLiveOutputBad(t *testing.T) {
	a := assert.New(t)

	cmd := NewCommand("/bin/cat /tmp/DOESNOTEXIST").SuppressStdout().SuppressStderr().Build()
	a.NotNil(cmd)
	exitCode, err := cmd.RunWithLiveOutput()
	a.Equal(1, exitCode)
	a.NotNil(err)
}
