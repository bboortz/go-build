package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

// ExecCommand executes a command with enabled error checks.
func ExecCommand(command string) (int, string, string) {
	return ExecCommandAllParams(command, true)
}

// ExecCommandWithoutErrCheck executes a command with disabled error checks.
func ExecCommandWithoutErrCheck(command string) (int, string, string) {
	return ExecCommandAllParams(command, false)
}

// ExecCommandAllParams executes a command via the os interface which let you define all possible parameter.
func ExecCommandAllParams(command string, checkError bool) (int, string, string) {

	// run command
	log.Debug("CMD: " + command)
	cmd := exec.Command("/bin/sh", "-c", command)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err := cmd.Run()

	// convert buffer to string
	stdoutStr := stdoutBuf.String()
	stderrStr := stderrBuf.String()

	// retrieve exit code
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	// logging
	if exitCode == 0 && err == nil {
		if stdoutStr != "" {
			log.Trace(stdoutStr)
		}
	} else {
		if stdoutStr != "" {
			log.Error(stdoutStr)
		}
	}
	if stderrStr != "" {
		log.Error(stderrStr)
	}

	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))
	if checkError && (err != nil || exitCode != 0) {
		log.Fatal(err)
	}

	return exitCode, stdoutStr, stderrStr
}
