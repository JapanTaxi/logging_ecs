package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Time    time.Time   `json:"time"`
	Error   string      `json:"error_text,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func Debug(message string) {
	outputLog("DEBUG", message, nil, nil)
}

func DebugWithDetail(message string, detail interface{}) {
	outputLog("DEBUG", message, nil, detail)
}

func Info(message string) {
	outputLog("INFO", message, nil, nil)
}

func InfoWithDetail(message string, detail interface{}) {
	outputLog("INFO", message, nil, detail)
}

func Warn(message string) {
	outputLog("WARN", message, nil, nil)
}

func WarnWithDetail(message string, detail interface{}) {
	outputLog("WARN", message, nil, detail)
}

func Error(message string, err error) {
	outputLog("ERROR", message, err, nil)
}

func ErrorWithDetail(message string, err error, detail interface{}) {
	outputLog("ERROR", message, err, detail)
}

func outputLog(level string, message string, err error, detail interface{}) {
	log := &Log{
		Level:   level,
		Message: message,
		Time:    GetJpTimeNow(),
		Detail:  detail,
	}

	if err != nil {
		log.Error = err.Error()
	}

	output(log)
}

func output(log interface{}) {
	logJSON, err := json.Marshal(log)
	if err == nil {
		fmt.Println(string(logJSON))
	} else {
		fmt.Println(err.Error())
	}
}
