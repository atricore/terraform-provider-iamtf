package cli

import "log"

type (
	Logger interface {
		Logger() interface{}

		Tracef(format string, v ...interface{})

		Trace(msg string)

		Debugf(format string, v ...interface{})

		Debug(msg string)

		Infof(format string, v ...interface{})

		Info(msg string)

		Warnf(format string, v ...interface{})

		Warn(msg string)

		Errorf(format string, v ...interface{})

		Error(msg string)
	}

	// DefaultLogger wraps standard log
	DefaultLogger struct {
		debug bool
	}
)

func NewDefaultLogger(debug bool) DefaultLogger {
	return DefaultLogger{debug: debug}
}

// Returns itself
func (l DefaultLogger) Logger() interface{} {
	return l
}

func (l DefaultLogger) Trace(msg string) {
	if l.debug {
		log.Print(msg)
	}
}

func (l DefaultLogger) Tracef(format string, v ...interface{}) {
	if l.debug {
		log.Printf(format, v...)
	}
}

func (l DefaultLogger) Debug(msg string) {
	if l.debug {
		log.Print(msg)
	}
}

func (l DefaultLogger) Debugf(format string, v ...interface{}) {
	if l.debug {
		log.Printf(format, v...)
	}
}

func (l DefaultLogger) Info(msg string) {
	if l.debug {
		log.Print(msg)
	}
}

func (l DefaultLogger) Infof(format string, v ...interface{}) {
	if l.debug {
		log.Printf(format, v...)
	}
}

func (l DefaultLogger) Warn(msg string) {
	if l.debug {
		log.Print(msg)
	}
}

func (l DefaultLogger) Warnf(format string, v ...interface{}) {
	if l.debug {
		log.Printf(format, v...)
	}
}

func (l DefaultLogger) Error(msg string) {
	if l.debug {
		log.Print(msg)
	}
}

func (l DefaultLogger) Errorf(format string, v ...interface{}) {
	if l.debug {
		log.Printf(format, v...)
	}
}
