package stringutil

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func init() {
}

func TestCheckEmptyGood(t *testing.T) {
	a := assert.New(t)
	if os.Getenv("TESTRUN") == "1" {
		CheckEmpty("testkey", "testvalue")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckEmptyGood")
	cmd.Env = append(os.Environ(), "TESTRUN=1")
	err := cmd.Run()
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	a.Nil(err)
	a.Equal(0, exitCode)
}

func TestCheckEmptyBad(t *testing.T) {
	a := assert.New(t)
	if os.Getenv("TESTRUN") == "1" {
		CheckEmpty("testkey", "")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckEmptyBad")
	cmd.Env = append(os.Environ(), "TESTRUN=1")
	err := cmd.Run()
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	a.NotNil(err)
	a.Equal(1, exitCode)
}
