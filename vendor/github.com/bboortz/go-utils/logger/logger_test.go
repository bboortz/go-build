package logger

import (
	//	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func logIndirect2(args ...interface{}) {
	log := NewLogger().Build()
	log.Critical(args...)
	log.Error(args...)
	log.Info(args...)
	log.Debug(args...)
	log.Trace(args...)
}
func logIndirect(args ...interface{}) {
	logIndirect2(args...)
}

func TestSetLevel(t *testing.T) {
	a := assert.New(t)
	log := NewLogger().Build()
	log.SetLevel(DEBUG)
	level := log.GetLevel()
	a.Equal(level, DEBUG)
}

func TestGetLevel(t *testing.T) {
	a := assert.New(t)
	log := NewLogger().Build()
	level := log.GetLevel()
	a.Equal(level, INFO)
}

func TestDebug(t *testing.T) {
	log := NewLogger().Build()
	log.Debug("test")
	log.Debug("test", "1")
	log.Debug("test", "1", 2)
	log.Debug("test", "1", 2, 3)
	log.Debug("test", "1", 2, 3, 4)
	log.Debug("test" + " foobar")
	log.Debug("test"+" foobar", 1)
}

func TestLogIndirect(t *testing.T) {
	logIndirect("test indirect")
}
