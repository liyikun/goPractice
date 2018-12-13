package logger

import (
	"log"
	"os"
)

var (
	Trace *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("error.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)

	
}
