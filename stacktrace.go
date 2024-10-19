package apierrors

import (
    "runtime"
    "fmt"
    "strings"
)

type stackTrace struct {
    Frames []stackFrame `json:"frames"`
}

type stackFrame struct {
    Function string `json:"function"`
    File     string `json:"file"`
    Line     int    `json:"line"`
}

func collectStackTrace() *stackTrace {
    var pcs [32]uintptr
    n := runtime.Callers(3, pcs[:])
    var frames []stackFrame
    for _, pc := range pcs[:n] {
        fn := runtime.FuncForPC(pc)
        if fn != nil {
            file, line := fn.FileLine(pc)
            frames = append(frames, stackFrame{
                Function: fn.Name(),
                File:     file,
                Line:     line,
            })
        }
    }
    return &stackTrace{Frames: frames}
}

// Improved String method for more readable stack traces.
func (st *stackTrace) String() string {
    var sb strings.Builder
    sb.WriteString("Stack Trace:\n")
    sb.WriteString("---------------------------------------------------\n")
    for i, frame := range st.Frames {
        sb.WriteString(fmt.Sprintf("#%d - Function: %s\n", i+1, frame.Function))
        sb.WriteString(fmt.Sprintf("     Location: %s:%d\n", frame.File, frame.Line))
        sb.WriteString("---------------------------------------------------\n")
    }
    return sb.String()
}

func (st *stackTrace) enhanceWithCause(original *stackTrace) {
    st.Frames = append(st.Frames, original.Frames...)
}
