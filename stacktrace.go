package apierrors

import (
    "runtime"
    "fmt"
)

type stack struct {
    Frames []string `json:"frames"`
}

func captureStack() *stack {
    var pcs [32]uintptr
    n := runtime.Callers(2, pcs[:])
    var frames []string
    for _, pc := range pcs[:n] {
        fn := runtime.FuncForPC(pc)
        if fn != nil {
            frames = append(frames, fmt.Sprintf("%s", fn.Name()))
        }
    }
    return &stack{Frames: frames}
}
