package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

type Logger struct {
	level int
	file  *os.File
}

var logFile Logger

func setLevel(level int) {
	logFile.level = level
}

func formatTime(t *time.Time) string {
	return t.Format(time.RFC3339)
}

func getwd() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filePathStr := filepath.Join(path, "logs")
	err = os.MkdirAll(filePathStr, 0755)
	if err != nil {
		return "", err
	}
	return filePathStr, nil
}

var loggerf *log.Logger

func Init(path, logName string, level int) error {
	if path == "" {
		filePath, _ := getwd()
		filepathStr := filepath.Join(filePath, logName)
		logfile, err := os.Create(filepathStr)
		if err != nil {
			return err
		}
		logFile.file = logfile
		setLevel(level)
	} else if logName == "" {
		filepathStr := path
		err := os.MkdirAll(filepathStr, 0644)
		if err != nil {
			return err
		}
		logName = filepath.Join(filepathStr, "logs.log")
		logfile, err := os.Create(logName)
		if err != nil {
			return err
		}
		logFile.file = logfile
	} else {
		filepathstr := path
		err := os.MkdirAll(filepathstr, 0644)
		if err != nil {
			return err
		}
		logName = filepath.Join(filepathstr, logName)
		logfile, err := os.Create(logName)
		if err != nil {
			return errors.New(fmt.Sprintln("Create File Error:", err))
		}
		logFile.file = logfile
		setLevel(level)
	}

	loggerf = log.New(logFile.file, "", log.Llongfile|log.Ltime|log.Ldate)
	return nil
}

func init() {
	logFile.level = LevelDebug
	loggerf = log.New(os.Stdout, "", log.Llongfile|log.Ltime|log.Ldate)
}

func Debug(format string, v ...interface{}) {
	if logFile.level >= LevelDebug {
		loggerf.Printf("[Debug] "+format, v...)
	}
}
