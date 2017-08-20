package utils

import (
	"bufio"
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

// ExecCommandWithOutput executes a command via the os interface which let you define all possible parameter.
func ExecCommandWithOutput(command string, withStdout bool, checkError bool) (int, error) {

	// run command
	log.Debug("CMD: " + command)
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	var err error
	stdoutDone := make(chan interface{})
	stderrDone := make(chan interface{})

	// start process via command
	if err = cmd.Start(); err != nil {
		return -1, err
	}

	// logging go routines
	go func() {
		if withStdout {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				log.Info(scanner.Text()) // Println will add back the final '\n'
			}
		}
		close(stdoutDone)
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Error(scanner.Text()) // Println will add back the final '\n'
		}
		close(stderrDone)
	}()

	// wait until commend has finished
	err = cmd.Wait()
	<-stdoutDone
	<-stderrDone

	// retrieve exit code
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))
	if checkError {
		if err != nil {
			log.Fatal(err)
		} else if exitCode != 0 {
			log.Fatal(fmt.Sprintf("EXIT CODE: %d", exitCode))
		}
	}

	return exitCode, err
}
