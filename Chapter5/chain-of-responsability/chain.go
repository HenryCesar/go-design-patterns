package chainofresponsability

import "io"

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {}

type SecondLogger struct {
	NextChain ChainLogger
}

func (f *SecondLogger) Next(s string) {}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {}
