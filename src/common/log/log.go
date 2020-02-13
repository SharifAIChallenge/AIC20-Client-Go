package log

import "fmt"

const (
	Verbose = iota
	Debug
	Info
	Warn
	Error
)

var Levels = [...]string{"Verbose", "Debug", "Info", "Warning", "Error"}

const DevMode bool = false
const LogLevel int = Warn

func V(tag, msg string) {
	log(Verbose, tag, msg)
}

func D(tag, msg string) {
	log(Debug, tag, msg)
}

func I(tag, msg string) {
	log(Info, tag, msg)
}

func W(tag, msg string) {
	log(Warn, tag, msg)
}

func E(tag, msg string) {
	log(Error, tag, msg)
}

func log(priority int, tag string, msg string) {
	if priority < LogLevel {
		return
	}
	fmt.Errorf("%s : %s - %s", Levels[priority], tag, msg)
}
