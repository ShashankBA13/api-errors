package apierrors

import (
	"fmt"
	"runtime"
)

// stackTrace holds detailed information about the call stack.
type stackTrace struct {
	Frames []string `json:"frames"`
}

// collectStackTrace captures the current call stack, recording the function name, file, and line number.
func collectStackTrace() *stackTrace {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // Skip 3 frames to get to the caller.
	var frames []string

	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		frames = append(frames, fmt.Sprintf("at %s\n\t%s:%d", fn.Name(), file, line))
	}

	return &stackTrace{Frames: frames}
}

// enhanceWithCause adds the original stack trace to the current one.
func (st *stackTrace) enhanceWithCause(cause *stackTrace) {
	if cause != nil {
		st.Frames = append(st.Frames, "caused by:")
		st.Frames = append(st.Frames, cause.Frames...)
	}
}
