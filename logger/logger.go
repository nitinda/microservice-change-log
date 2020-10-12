package logger

import (
	"io"
	"log"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {

	Trace = log.New(traceHandle, "change-log-api : TRACE : ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "change-log-api : INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "change-log-api : WARNING : ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "change-log-api : ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
}
