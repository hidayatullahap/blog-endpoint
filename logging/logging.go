package logging

import (
	"io"
	"log"
)

type Logger struct {
	*log.Logger
}

var (
	Trace   Logger
	Info    Logger
	Warning Logger
	Error   Logger
)

func Initialize(traceWriter io.Writer, infoWriter io.Writer, warningWriter io.Writer, errorWriter io.Writer) {
	Trace.Logger = log.New(traceWriter, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info.Logger = log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning.Logger = log.New(warningWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error.Logger = log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l Logger) Log(message string) {
	l.Logger.Println(message)
}
