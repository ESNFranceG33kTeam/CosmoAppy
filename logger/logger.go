package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	timestamp time.Time
	level     string
	endpoint  string
	message   string
	detailes  error
}

type Loggers []Logger

const InfoColor = "\033[0;32m"
const WarningColor = "\033[0;33m"
const ErrorColor = "\033[1;31m"
const CriticalColor = "\033[0;31m"
const NoneColor = "\033[0m"

func LogInfo(endpoint string, message string) {
	var mylog Logger

	mylog.timestamp = time.Now()
	mylog.level = "INFO"
	mylog.endpoint = endpoint
	mylog.message = message

	fmt.Fprintln(os.Stdout, InfoColor, "timestamp=\""+mylog.timestamp.Format("2006-01-02 15:04:05")+"\" level="+mylog.level+" endpoint="+mylog.endpoint+" message=\""+mylog.message+"\"", NoneColor)
}

func LogWarning(endpoint string, message string, detailes error) {
	var mylog Logger

	mylog.timestamp = time.Now()
	mylog.level = "WARNING"
	mylog.endpoint = endpoint
	mylog.message = message
	mylog.detailes = detailes

	fmt.Fprintln(os.Stderr, WarningColor, "timestamp=\""+mylog.timestamp.Format("2006-01-02 15:04:05")+"\" level="+mylog.level+" endpoint="+mylog.endpoint+" message=\""+mylog.message+"\" detailes=\""+mylog.detailes.Error()+"\"", NoneColor)
}

func LogError(endpoint string, message string, detailes error) {
	var mylog Logger

	mylog.timestamp = time.Now()
	mylog.level = "ERROR"
	mylog.endpoint = endpoint
	mylog.message = message
	mylog.detailes = detailes

	fmt.Fprintln(os.Stderr, ErrorColor, "timestamp=\""+mylog.timestamp.Format("2006-01-02 15:04:05")+"\" level="+mylog.level+" endpoint="+mylog.endpoint+" message=\""+mylog.message+"\" detailes=\""+mylog.detailes.Error()+"\"", NoneColor)
}

func LogCritical(endpoint string, message string, detailes error, exit ...bool) {
	var mylog Logger

	mylog.timestamp = time.Now()
	mylog.level = "CRITICAL"
	mylog.endpoint = endpoint
	mylog.message = message
	mylog.detailes = detailes

	fmt.Fprintln(os.Stderr, CriticalColor, "timestamp=\""+mylog.timestamp.Format("2006-01-02 15:04:05")+"\" level="+mylog.level+" endpoint="+mylog.endpoint+" message=\""+mylog.message+"\" detailes=\""+mylog.detailes.Error()+"\"", NoneColor)
	if len(exit) == 0 {
		os.Exit(1)
	}
}
