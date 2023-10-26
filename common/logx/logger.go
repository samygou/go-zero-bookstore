package logx

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	l          *logrus.Logger
	callerSkip int
}

func New() *Logger {
	return &Logger{l: logrus.New(), callerSkip: 2}
}

var errMissingValue = errors.New("(MISSING)")

var Std = New()
var inner = New()

func (l *Logger) SetCallerSkip(skip int) {
	l.callerSkip = skip
}

func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.l.SetFormatter(formatter)
}

func (l *Logger) SetLevel(level Level) {
	l.l.SetLevel(level)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) Panic(args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) Fatal(args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Errorf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.l.WithFields(logrus.Fields{
		"internal-caller": l.getCaller(),
	}).Fatalf(template, args...)
}

func (l *Logger) WithFields(fields Fields) *Entry {
	fields["internal-caller"] = l.getCaller()
	return l.l.WithFields(fields)
}

func (l *Logger) ErrStack(err error) {
	l.l.Error(err)
	l.l.Error(string(debug.Stack()))
	//debug.PrintStack()
}

func (l *Logger) Log(keyvals ...interface{}) error {
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		} else {
			fields[fmt.Sprint(keyvals[i])] = errMissingValue
		}
	}
	l.WithFields(fields).Info()
	return nil
}

func (l *Logger) Print(level Level, caller string, args ...interface{}) {
	switch level {
	case PanicLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Panic(args...)
	case FatalLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Fatal(args...)
	case ErrorLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Error(args...)
	case WarnLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Warn(args...)
	case InfoLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Info(args...)
	case DebugLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Debug(args...)
	case TraceLevel:
		l.l.WithFields(logrus.Fields{
			"internal-caller": caller,
		}).Trace(args...)
	}
}
func (l *Logger) getCaller() string {
	dep := 3
	_, file, line, _ := runtime.Caller(l.callerSkip)
	fnodes := strings.Split(file, "/")
	caller := ":" + strconv.Itoa(line)
	for i := 0; i < len(fnodes) && i < dep; i++ {
		if i > 0 {
			caller = "/" + caller
		}
		caller = fnodes[len(fnodes)-(i+1)] + caller
	}

	caller = strings.Replace(caller, ".go", "", -1)

	return caller
}
