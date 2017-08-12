package logger

import (
	"fmt"
	"github.com/bboortz/go-utils/slices"
	"github.com/bboortz/go-utils/stack"
	"log"
	"os"
	"time"
)

/*
 * Log levels
 */
type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
	TRACE
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"TRACE",
}

/*
 * Interface Definition
 */
type Logger interface {
	SetLevel(Level)
	SetLevelWithStr(string)
	GetLevel() Level
	Log(level Level, caller string, args ...interface{})
	Fatal(args ...interface{})
	Critical(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Notice(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
}

type LoggerBuilder interface {
	SetLevel(Level) LoggerBuilder
	Build() Logger
}

type loggerBuilder struct {
	level Level
}

/*
 * Builder Methods
 */
func NewLogger() LoggerBuilder {
	return &loggerBuilder{level: INFO}
}

func (b *loggerBuilder) SetLevel(level Level) LoggerBuilder {
	b.level = level
	return b
}

func (b *loggerBuilder) SetLevelWithStr(levelStr string) LoggerBuilder {
	level := Level(slices.IndexOf(levelStr, levelNames))
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
	level := Level(slices.IndexOf(levelStr, levelNames))
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

func (l *logger) Warning(args ...interface{}) {
	if l.level < WARNING {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(WARNING, callerStr, args)
}

func (l *logger) Notice(args ...interface{}) {
	if l.level < NOTICE {
		return
	}
	callerStr := stack.GetCallingMethodName()
	l.Log(NOTICE, callerStr, args)
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
