package main

import (
	"os"
)

type Output struct {
	Writer *os.File
}

func NewOutput(w *os.File) Output {
	return Output{Writer: w}
}

func (o *Output) WriteString(s string) {
	o.Writer.Write([]byte(s))
	o.Writer.Sync()
}

func (o *Output) Clear() {
	o.Writer.Write([]byte("\033[2J"))
	o.Writer.Write([]byte("\033[H"))
	o.Writer.Sync()
}
