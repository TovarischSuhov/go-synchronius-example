package log

import (
	"fmt"
	"time"
)

type logLevel int

const (
	Debug logLevel = iota
	Info
	Warn
	Error
	Fatal
)

var colors = map[logLevel]string{
	Debug: "",
	Info:  "",
	Warn:  "",
	Error: "",
	Fatal: "",
}

const finishColor = ""

var (
	UseColors = false
	LogLevel  = Info
)

func printMsg(level logLevel, msg string) {
	ts := time.Now().Format(time.RFC3339)
	prefix := "[" + level.String() + "]"
	if UseColors {
		prefix = colors[level] + prefix + finishColor
	}
	fmt.Printf("%s %s %s", ts, prefix, msg)
}

func Debug(msg string) {
	printMsg(Debug, msg)
}

func Info(msg string) {
	printMsg(Info, msg)
}

func Warn(msg string) {
	printMsg(Warn, msg)
}

func Error(msg string) {
	printMsg(Error, msg)
}

func Panic(msg string) {
	printMsg(Panic, msg)
	panic("sends panic")
}
