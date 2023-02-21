// Package log - yet another logger, сохранён интерфейс из исходного примера. В реальных проектах будет что-то
// типа [*zerolog.Log|*logrus.Logger|...] без интерфейсов, или свои обёртки. Может быть, добавление slog в стандартную
// ("ыуу https://github.com/golang/go/issues/56345) библиотеку скоро позволит избавиться от этого разнообразия
package log

import (
	"fmt"
	"log"
)

type Log struct {
	logger *log.Logger
}

func NewLog(opts ...Option) *Log {
	o := NewDefaultOptions()
	o.apply(opts...)

	return &Log{
		logger: log.New(o.out, o.prefix, log.LstdFlags),
	}
}

func (l *Log) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v)
	l.logger.Printf("[Error]: %s\n", msg)
}

func (l *Log) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v)
	l.logger.Printf("[Info]: %s\n", msg)
}
