package logger

import (
	"fmt"
	"github.com/bboortz/go-utils/sliceutil"
	"github.com/bboortz/go-utils/stack"
	"log"
	"os"
	"time"
)

// Level is the type for log levels
type Level int

const (
	// CRITICAL is the log level for critical situations
	CRITICAL Level = iota
	// ERROR is the log level for errors
	ERROR
	// INFO is the log level for information
	INFO
	// DEBUG is the log level for debug messages
	DEBUG
	// TRACE is the log level for traces
	TRACE
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"INFO",
	"DEBUG",
	"TRACE",
}

// Logger is the main interface for this logger
type Logger interface {
	SetLevel(Level)
	SetLevelWithStr(string)
	GetLevel() Level
	Log(level Level, caller string, args ...interface{})
	Fatal(args ...interface{})
	Critical(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
}

// Builder is the interface to build a new logger
type Builder interface {
	SetLevel(Level) Builder
	Build() Logger
}

type loggerBuilder struct {
	level Level
}

// NewLogger is the create function for the Builder
func NewLogger() Builder {
	return &loggerBuilder{level: INFO}
}

func (b *loggerBuilder) SetLevel(level Level) Builder {
	b.level = level
	return b
}

func (b *loggerBuilder) SetLevelWithStr(levelStr string) Builder {
	level := Level(sliceutil.IndexOf(levelStr, levelNames))
	b.SetLevel(level)
	return b
}

func (b *loggerBuilder) Build() Logger {
	return &logger{
		logStream: log.New(os.Stdout, "", 0),
		level:     b.level,
	}
}

/*
 * OOP Methods
 */

type logger struct {
	logStream *log.Logger
	level     Level
}

func (l *logger) SetLevel(level Level) {
	l.level = level
}

func (l *logger) SetLevelWithStr(levelStr string) {
	level := Level(sliceutil.IndexOf(levelStr, levelNames))
	l.SetLevel(level)
}

func (l *logger) GetLevel() Level {
	return l.level
}

func (l *logger) Log(level Level, caller string, args ...interface{}) {
	currentStr := time.Now().Format(time.RFC3339)
	levelStr := levelNames[level]
	msg := fmt.Sprintln(args...)
	logLine := fmt.Sprintf("%s %s > %-4s %s", currentStr, caller, levelStr, msg)
	l.logStream.Printf(logLine)
}

func (l *logger) Fatal(args ...interface{}) {
	if l.level < CRITICAL {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(CRITICAL, callerStr, args)
	os.Exit(1)
}

func (l *logger) Critical(args ...interface{}) {
	if l.level < CRITICAL {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(CRITICAL, callerStr, args)
}

func (l *logger) Error(args ...interface{}) {
	if l.level < ERROR {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(ERROR, callerStr, args)
}

func (l *logger) Info(args ...interface{}) {
	if l.level < INFO {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(INFO, callerStr, args)
}

func (l *logger) Debug(args ...interface{}) {
	if l.level < DEBUG {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(DEBUG, callerStr, args)
}

func (l *logger) Trace(args ...interface{}) {
	if l.level < TRACE {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(TRACE, callerStr, args)
}
