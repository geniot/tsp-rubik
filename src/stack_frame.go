package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
)

// A StackFrame contains all necessary information about to generate a line
// in a callstack.
type StackFrame struct {
	File           string
	LineNumber     int
	Name           string
	Package        string
	ProgramCounter uintptr
}

// newStackFrame populates a stack frame object from the program counter.
func newStackFrame(pc uintptr) (frame StackFrame) {

	frame = StackFrame{ProgramCounter: pc}
	if frame.Func() == nil {
		return
	}
	frame.Package, frame.Name = packageAndName(frame.Func())

	// pc -1 because the program counters we use are usually return addresses,
	// and we want to show the line that corresponds to the function call
	frame.File, frame.LineNumber = frame.Func().FileLine(pc - 1)
	return

}

// Func returns the function that this stackframe corresponds to
func (frame *StackFrame) Func() *runtime.Func {
	if frame.ProgramCounter == 0 {
		return nil
	}
	return runtime.FuncForPC(frame.ProgramCounter)
}

// String returns the stackframe formatted in the same way as go does
// in runtime/debug.Stack()
func (frame *StackFrame) String() string {
	str := fmt.Sprintf("%s:%d (0x%x)\n", frame.File, frame.LineNumber, frame.ProgramCounter)

	source, err := frame.SourceLine()
	if err != nil {
		return str
	}

	return str + fmt.Sprintf("\t%s: %s\n", frame.Name, source)
}

// SourceLine gets the line of code (from File and Line) of the original source if possible
func (frame *StackFrame) SourceLine() (string, error) {
	data, err := ioutil.ReadFile(frame.File)

	if err != nil {
		return "", err
	}

	lines := bytes.Split(data, []byte{'\n'})
	if frame.LineNumber <= 0 || frame.LineNumber >= len(lines) {
		return "???", nil
	}
	// -1 because line-numbers are 1 based, but our array is 0 based
	return string(bytes.Trim(lines[frame.LineNumber-1], " \t")), nil
}
