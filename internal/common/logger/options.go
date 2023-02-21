package log

import (
	"io"
	"os"
)

type LoggerOptions struct {
	out    io.Writer
	level  string
	prefix string
}

type Option func(options *LoggerOptions)

func (o *LoggerOptions) apply(options ...Option) {
	for _, opt := range options {
		opt(o)
	}
}

func NewDefaultOptions() *LoggerOptions {
	return &LoggerOptions{
		level: "info",
		out:   os.Stdout,
	}
}

func WithLogLevel(level string) Option {
	return func(o *LoggerOptions) {
		o.level = level
	}
}

func WithPrefix(prefix string) Option {
	return func(o *LoggerOptions) {
		o.prefix = prefix
	}
}

func WithOut(out io.Writer) Option {
	return func(o *LoggerOptions) {
		o.out = out
	}
}
