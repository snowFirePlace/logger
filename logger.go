package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	fTime = "02.01.2006 15:04:05"
	std   = New()
)

//
// Log new
type Log struct {
	MU    sync.Mutex
	Out   io.Writer
	Level Level
	Buf   []byte
}
type Level struct {
	Number int
}

func New() *Log {
	return &Log{
		Out:   os.Stderr,
		Level: Level{Number: 3},
	}
}
func (l *Log) SetOutput(w io.Writer) {
	l.MU.Lock()
	defer l.MU.Unlock()
	l.Out = w
}
func (l *Log) SetLevel(le string) {
	switch le {
	case "Track":
		l.Level.Number = 1
		fTime = "02.01.2006 15:04:05.000000000"
	case "Debug":
		l.Level.Number = 2
		fTime = "02.01.2006 15:04:05.000000"
	case "Info":
		l.Level.Number = 3
	}
}
func (l *Log) out(s, m string) {
	l.MU.Lock()
	defer l.MU.Unlock()
	if len(strings.Split(m, "\r\n")) != 1 {
		m = fmt.Sprintf("%s\t| %s |\t----MESSAGE----\r\n%s\r\n", time.Now().Format(fTime), s, m)
	} else {
		m = fmt.Sprintf("%s\t| %s |\t%s\r\n", time.Now().Format(fTime), s, m)
	}
	l.Buf = l.Buf[:0]
	l.Buf = append(l.Buf, []byte(m)...)
	_, _ = l.Out.Write(l.Buf)
}

func (l *Log) Debug(args ...interface{}) {
	if l.Level.Number <= 2 {
		if len(args) >= 2 {
			a := fmt.Sprintln(args...)
			l.out("DEBUG", a[:len(a)-1])
			return
		}
		l.out("DEBUG", fmt.Sprint(args...))
	}
}

func (l *Log) Info(args ...interface{}) {
	if l.Level.Number <= 3 {
		if len(args) >= 2 {
			a := fmt.Sprintln(args...)
			l.out("INFO ", a[:len(a)-1])
			return
		}
		l.out("INFO ", fmt.Sprint(args...))

	}
}

func (l *Log) Error(args ...interface{}) {
	if len(args) >= 2 {
		a := fmt.Sprintln(args...)
		l.out("ERROR", a[:len(a)-1])
		return
	}
	l.out("ERROR", fmt.Sprint(args...))
}

func (l *Log) Fatal(args ...interface{}) {
	if len(args) >= 2 {
		a := fmt.Sprintln(args...)
		l.out("FATAL", a[:len(a)-1])
	} else {
		l.out("FATAL", fmt.Sprint(args...))
	}
	os.Exit(1)
}

func (l *Log) Fatalf(s string, args ...interface{}) {
	l.out("FATAL", fmt.Sprintf(s, args...))
	os.Exit(1)
}

func Debug(args ...interface{}) {
	if std.Level.Number <= 2 {
		if len(args) >= 2 {
			a := fmt.Sprintln(args...)
			std.out("DEBUG", a[:len(a)-1])
			return
		}
		std.out("DEBUG", fmt.Sprint(args...))
	}
}
func Info(args ...interface{}) {
	if std.Level.Number <= 3 {
		if len(args) >= 2 {
			a := fmt.Sprintln(args...)
			std.out("INFO ", a[:len(a)-1])
			return
		}
		std.out("INFO", fmt.Sprint(args...))
	}
}
func Error(args ...interface{}) {
	if len(args) >= 2 {
		a := fmt.Sprintln(args...)
		std.out("ERROR", a[:len(a)-1])
		return
	}
	std.out("ERROR", fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	if len(args) >= 2 {
		a := fmt.Sprintln(args...)
		std.out("FATAL", a[:len(a)-1])
	} else {
		std.out("FATAL", fmt.Sprint(args...))
	}
	os.Exit(1)
}

func Fatalf(s string, args ...interface{}) {
	std.out("FATAL", fmt.Sprintf(s, args...))
	os.Exit(1)
}
