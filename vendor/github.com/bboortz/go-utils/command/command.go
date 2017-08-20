package command

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/bboortz/go-utils/logger"
	"os/exec"
	"syscall"
)

var log = logger.NewLogger().Build()

// Command is the main interface
type Command interface {
	Run() (int, string, string, error)
	RunWithLiveOutput() (int, error)
}

type command struct {
	cmd        *exec.Cmd
	stdoutOn   bool
	stderrOn   bool
	checkError bool
}

func (s *command) logCommand() {
	log.Debug("CMD: " + s.cmd.Path)
}

// Run executes a command via the os interface and returns stdout and stderr output
func (s *command) Run() (int, string, string, error) {
	s.logCommand()

	var stdoutBuf, stderrBuf bytes.Buffer
	s.cmd.Stdout = &stdoutBuf
	s.cmd.Stderr = &stderrBuf
	var stdoutStr string
	var stderrStr string
	err := s.cmd.Run()

	// convert buffer to string
	if s.stdoutOn {
		stdoutStr = stdoutBuf.String()
	}
	if s.stderrOn {
		stderrStr = stderrBuf.String()
	}

	// retrieve exit code
	waitStatus := s.cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))
	if s.checkError && (err != nil || exitCode != 0) {
		log.Fatal(err)
	}

	return exitCode, stdoutStr, stderrStr, err
}

// RunWithLiveOutput executes a command via the os interface and produced live output
func (s *command) RunWithLiveOutput() (int, error) {
	s.logCommand()

	var err error
	stdoutDone := make(chan interface{})
	stderrDone := make(chan interface{})
	stdout, _ := s.cmd.StdoutPipe()
	stderr, _ := s.cmd.StderrPipe()

	// start process via command
	if err = s.cmd.Start(); err != nil {
		return -1, err
	}

	// logging go routines
	go func() {
		if s.stdoutOn {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				log.Info(scanner.Text()) // Println will add back the final '\n'
			}
		}
		close(stdoutDone)
	}()
	go func() {
		if s.stderrOn {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				log.Error(scanner.Text()) // Println will add back the final '\n'
			}
		}
		close(stderrDone)
	}()

	// wait until commend has finished
	err = s.cmd.Wait()
	<-stdoutDone
	<-stderrDone

	// retrieve exit code
	waitStatus := s.cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()

	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))
	if s.checkError && (err != nil || exitCode != 0) {
		log.Fatal(err)
	}

	return exitCode, err
}
