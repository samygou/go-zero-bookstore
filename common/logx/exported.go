package logx

import (
	"os"

	"github.com/sirupsen/logrus"
)

//
//var (
//	flagDebug = flag.Bool("log-debug", false, "log debug switch")
//)

type Fields = logrus.Fields
type Entry = logrus.Entry

func init() {
	Std.SetFormatter(&Formatter{})
	inner.SetFormatter(&Formatter{})
	inner.SetCallerSkip(3)

	level := os.Getenv("IMO_LOG_LEVEL")
	switch level {
	case "DEBUG":
		Std.SetLevel(logrus.DebugLevel)
		inner.SetLevel(logrus.DebugLevel)
	case "TRACE":
		Std.SetLevel(logrus.TraceLevel)
		inner.SetLevel(logrus.TraceLevel)
	case "INFO":
		Std.SetLevel(logrus.InfoLevel)
		inner.SetLevel(logrus.InfoLevel)
	case "WARN":
		Std.SetLevel(logrus.WarnLevel)
		inner.SetLevel(logrus.WarnLevel)
	case "ERROR":
		Std.SetLevel(logrus.ErrorLevel)
		inner.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		Std.SetLevel(logrus.FatalLevel)
		inner.SetLevel(logrus.FatalLevel)
	case "PANIC":
		Std.SetLevel(logrus.PanicLevel)
		inner.SetLevel(logrus.PanicLevel)
	default:
		Std.SetLevel(logrus.InfoLevel)
		inner.SetLevel(logrus.InfoLevel)
	}
}

type Level = logrus.Level

// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func SetLevel(level Level) {
	inner.SetLevel(level)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	inner.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	inner.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	inner.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	inner.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	inner.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	inner.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	inner.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	inner.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	inner.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	inner.Errorf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	inner.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	inner.Fatalf(template, args...)
}

func WithFields(fields Fields) *Entry {
	return inner.WithFields(fields)
}

func ErrStack(err error) {
	inner.ErrStack(err)
}

// go kit log
func Log(keyvals ...interface{}) error {
	return inner.Log(keyvals...)
}

func Print(level Level, caller string, args ...interface{}) {
	inner.Print(level, caller, args...)
}
