package iamtf

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

// ProviderLogger wraps hclog.Logger
type ProviderLogger struct {
	wrapped hclog.Logger
}

func (l ProviderLogger) Logger() interface{} {
	return l.wrapped
}

func (l ProviderLogger) Trace(msg string) {
	l.wrapped.Trace(msg)
}

func (l ProviderLogger) Tracef(format string, v ...interface{}) {
	l.wrapped.Trace(fmt.Sprintf(format, v...))
}
func (l ProviderLogger) Debug(msg string) {
	l.wrapped.Debug(msg)
}

func (l ProviderLogger) Debugf(format string, v ...interface{}) {
	l.wrapped.Debug(fmt.Sprintf(format, v...))
}

func (l ProviderLogger) Info(msg string) {
	l.wrapped.Info(msg)
}

func (l ProviderLogger) Infof(format string, v ...interface{}) {
	l.wrapped.Info(fmt.Sprintf(format, v...))
}

func (l ProviderLogger) Warn(msg string) {
	l.wrapped.Warn(msg)
}

func (l ProviderLogger) Warnf(format string, v ...interface{}) {
	l.wrapped.Warn(fmt.Sprintf(format, v...))
}

func (l ProviderLogger) Error(msg string) {
	l.wrapped.Error(msg)
}

func (l ProviderLogger) Errorf(format string, v ...interface{}) {
	l.wrapped.Error(fmt.Sprintf(format, v...))
}
