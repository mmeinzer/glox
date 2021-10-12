package report

import (
	"log"
	"os"
)

type ErrorReporter interface {
	Error(line int, message string)
}

type errorReporter struct {
	hadError bool
	logger   *log.Logger
}

func NewErrorReporter() *errorReporter {
	logger := log.New(os.Stderr, "", 0)
	return &errorReporter{logger: logger}
}

func (e *errorReporter) Error(line int, message string) {
	e.report(line, "", message)
}

func (e *errorReporter) report(line int, where string, message string) {
	e.hadError = true
	e.logger.Printf("[line %d] Error %s: %s\n", line, where, message)
}
