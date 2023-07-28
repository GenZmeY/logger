package logger

import (
	"io"
	"log"
	"os"
)

const (
	LL_None LogLevel = iota
	LL_Fatal
	LL_Error
	LL_Warning
	LL_Info
	LL_Debug
	LL_Trace
	LL_All
)

type LogLevel int

func (ll LogLevel) String() string {
	return [...]string{
		"None",
		"Fatal",
		"Error",
		"Warning",
		"Info",
		"Debug",
		"Trace",
		"All",
	}[ll]
}

type Logger struct {
	// "const" vars
	stdout *log.Logger
	stderr *log.Logger
	trace  *log.Logger

	lFatal   *log.Logger
	lError   *log.Logger
	lWarning *log.Logger
	lInfo    *log.Logger
	lDebug   *log.Logger
	lTrace   *log.Logger

	// customizable vars
	logLevel      LogLevel
	printLogLevel bool
}

func New(out io.Writer, tag string, flag int, logLevel LogLevel, printLogLevel bool) *Logger {
	var l Logger

	if out == nil {
		l.stdout = log.New(os.Stdout, tag, flag)
		l.stderr = log.New(os.Stderr, tag, flag)
		l.trace = log.New(os.Stdout, tag, flag|log.Llongfile)
	} else {
		l.stdout = log.New(out, tag, flag)
		l.stderr = l.stdout
		l.trace = log.New(out, tag, flag|log.Llongfile)
	}

	l.lFatal = l.stderr
	l.lError = l.stderr
	l.lWarning = l.stdout
	l.lInfo = l.stdout
	l.lDebug = l.stdout
	l.lTrace = l.trace

	l.logLevel = logLevel
	l.printLogLevel = printLogLevel

	return &l
}

func (l *Logger) format(msg string) string {
	if l.printLogLevel {
		msg = l.logLevel.String() + ": " + msg
	}
	return msg
}

func (l *Logger) Fatal(msg string, v ...interface{}) {
	if l.logLevel >= LL_Fatal {
		l.lFatal.Printf(l.format(msg), v...)
	}
}

func (l *Logger) Error(msg string, v ...interface{}) {
	if l.logLevel >= LL_Error {
		l.lError.Printf(l.format(msg), v...)
	}
}

func (l *Logger) Warning(msg string, v ...interface{}) {
	if l.logLevel >= LL_Warning {
		l.lWarning.Printf(l.format(msg), v...)
	}
}

func (l *Logger) Info(msg string, v ...interface{}) {
	if l.logLevel >= LL_Info {
		l.lInfo.Printf(l.format(msg), v...)
	}
}

func (l *Logger) Debug(msg string, v ...interface{}) {
	if l.logLevel >= LL_Debug {
		l.lDebug.Printf(l.format(msg), v...)
	}
}

func (l *Logger) Trace(msg string, v ...interface{}) {
	if l.logLevel >= LL_Trace {
		l.lTrace.Printf(l.format(msg), v...)
	}
}

var def = New(nil, "", 0, LL_All, true)

func Default() *Logger { return def }

func Fatal(msg string, v ...interface{}) {
	def.Fatal(msg, v...)
}

func Error(msg string, v ...interface{}) {
	def.Error(msg, v...)
}

func Warning(msg string, v ...interface{}) {
	def.Warning(msg, v...)
}

func Info(msg string, v ...interface{}) {
	def.Info(msg, v...)
}

func Debug(msg string, v ...interface{}) {
	def.Debug(msg, v...)
}

func Trace(msg string, v ...interface{}) {
	def.Trace(msg, v...)
}
