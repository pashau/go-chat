package trace

import (
	"fmt"
	"io"
)

/* ######## Tracer ######## */

// interface - describes the methods Tracer objects will implement
type Tracer interface {
	Trace(...interface{})
}

// tracer implementation: the 'tracer type' has an 'io.Writer' field called 'out' which is where we will write the trace output to
type tracer struct {
	out io.Writer
}

// the 'Trace' method exactly matches the method required by the 'Tracer' interface, although it doesn't do anything yet.
func (t *tracer) Trace(a ...interface{}) {
	//format & write the trace details to the 'out' writer
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// method-creates a new instance of a Tracer.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

/* ######## nilTracer ######## */

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// method-gets a Tracer that does nothing.
func Off() Tracer {
	return &nilTracer{}
}
