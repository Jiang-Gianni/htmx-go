package mylogger

import (
	"io"
	"log"
)

const (
	INFO  = "INFO "
	ERROR = "ERROR "
	DEBUG = "DEBUG "
	WARN  = "WARN "
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
	Warn  *log.Logger
}

func New(w io.Writer, flag int) *Logger {
	return &Logger{
		Info:  log.New(w, INFO, flag),
		Error: log.New(w, ERROR, flag),
		Debug: log.New(w, DEBUG, flag),
		Warn:  log.New(w, WARN, flag),
	}
}
