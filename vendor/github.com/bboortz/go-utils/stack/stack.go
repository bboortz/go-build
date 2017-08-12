package stack

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/*
 * prints the stack to stderr
 */
func PrintStack() {
	os.Stderr.Write(GetStack())
}

/*
 * return a stack trace
 * Stack returns a formatted stack trace of the goroutine that calls it.
 * It calls runtime.Stack with a large enough buffer to capture the entire trace.
 */
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

func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func GetCurrentMethodName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	arr := strings.Split(f.Name(), ".")
	return arr[len(arr)-1]
}

func GetCallingMethodName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	arr := strings.Split(f.Name(), ".")
	return arr[len(arr)-1]
}
