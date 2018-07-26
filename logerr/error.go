package logerr

import (
	"runtime"
	"time"
)

// Err wraps the standard go error package with lots of convenience stuff
type Errlog struct {
	message string // Should only contain text if there is an error
	when    time.Time
	data    []string
	stack   []StackFrame
	err     *error // options error causing this error
}

type StackFrame struct {
	File string
	Line int
	PC   int // program counter
}

var (
	ErrDefault Errlog
	ErrNil     Errlog
	ErrNotNil  Errlog
	ErrEmpty   Errlog
	ErrError   Errlog
)

func init() {
	ErrDefault = Errlog{
		message: "",
		when:    time.Now(),
		stack:   nil,
		data:    nil,
		err:     nil,
	}
}

func GetParentFrame(skip int) *StackFrame {
	var sf *StackFrame = nil

	pc, fname, line, ok := runtime.Caller(1) // skip our caller?
	if ok {
		sf = &StackFrame{fname, line, int(pc)}
	}
	return sf
}

// ErrError will create an error when the para err != nil
func NewError(message string) *Errlog {
	if message == "" {
		return nil
	}
	pf := GetParentFrame(2)
	e := ErrDefault
	e.message = message
	e.stack = append(e.stack, *pf)
	return &e
}

// Error returns the error string and satisfies the go package error interface
func (err *Errlog) Error() string {
	return err.message
}

// When returns the timestamp
func (err *Errlog) When() time.Time {
	return err.when
}

// Append appends data to the data member
func (err *Errlog) AppendData(data string) {
	err.data = append(err.data, data)
}

// Push the given file name and line on the call stack
func (err *Errlog) PushStack(fname string, lineno, pc int) {
	fl := &StackFrame{fname, lineno, pc}
	err.stack = append(err.stack, *fl)
}

// CallStack returns the corresponding structure
func (err *Errlog) CallStack() []StackFrame {
	return err.stack
}

// Data return the data we may have collected along the stack
func (err *Errlog) Data() []string {
	return err.data
}
