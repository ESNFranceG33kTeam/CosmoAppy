package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
)

type Logger struct {
	loggerLevel int
	drawing     string
	timestamp   time.Time
	level       string
	endpoint    string
	message     string
	detailes    error
}

const DrawColor = "\033[96;1m"
const InfoColor = "\033[0;32m"
const WarningColor = "\033[0;33m"
const ErrorColor = "\033[1;31m"
const CriticalColor = "\033[0;31m"
const NoneColor = "\033[0m"

var mylog Logger

func (log *Logger) LogInit() {
	log.loggerLevel = helpers.TheAppConfig().Loglevel
}

func (log *Logger) LogDraw(drawing string) {
	log.drawing = drawing

	fmt.Fprintln(os.Stdout, DrawColor, string(log.drawing), NoneColor)
}

func (log *Logger) LogInfo(endpoint string, message string) {
	log.timestamp = time.Now()
	log.level = "INFO"
	log.endpoint = endpoint
	log.message = message

	if log.loggerLevel == 0 {
		fmt.Fprintln(
			os.Stdout,
			InfoColor,
			"timestamp=\""+log.timestamp.Format(
				"2006-01-02 15:04:05",
			)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+"\"",
			NoneColor,
		)
	}
}

func (log *Logger) LogWarning(endpoint string, message string, detailes error) {
	log.timestamp = time.Now()
	log.level = "WARNING"
	log.endpoint = endpoint
	log.message = message
	log.detailes = detailes

	if log.loggerLevel <= 1 {
		fmt.Fprintln(
			os.Stderr,
			WarningColor,
			"timestamp=\""+log.timestamp.Format(
				"2006-01-02 15:04:05",
			)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+
				"\" detailes=\""+log.detailes.Error()+"\"",
			NoneColor,
		)

	}
}

func (log *Logger) LogError(endpoint string, message string, detailes error) {
	log.timestamp = time.Now()
	log.level = "ERROR"
	log.endpoint = endpoint
	log.message = message
	log.detailes = detailes

	fmt.Fprintln(
		os.Stderr,
		ErrorColor,
		"timestamp=\""+log.timestamp.Format(
			"2006-01-02 15:04:05",
		)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+
			"\" detailes=\""+log.detailes.Error()+"\"",
		NoneColor,
	)
}

func (log *Logger) LogCritical(endpoint string, message string, detailes error, exit ...bool) {
	log.timestamp = time.Now()
	log.level = "CRITICAL"
	log.endpoint = endpoint
	log.message = message
	log.detailes = detailes

	fmt.Fprintln(
		os.Stderr,
		CriticalColor,
		"timestamp=\""+log.timestamp.Format(
			"2006-01-02 15:04:05",
		)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+
			"\" detailes=\""+log.detailes.Error()+"\"",
		NoneColor,
	)
	if len(exit) == 0 {
		os.Exit(1)
	}
}

func GetLogger() *Logger {
	return &mylog
}
