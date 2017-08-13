package stack

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// PrintStack prints the current call stack
func PrintStack() {
	_, _ = os.Stderr.Write(GetStack())
}

// GetStack retrieves a formatted call stack
func GetStack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

// Trace retrieves the current point of the call stack
func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

// GetCurrentMethodName retrieves the current method name
func GetCurrentMethodName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	arr := strings.Split(f.Name(), ".")
	return arr[len(arr)-1]
}

// GetCallingMethodName retrieves the calling method name
func GetCallingMethodName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	arr := strings.Split(f.Name(), ".")
	return arr[len(arr)-1]
}
