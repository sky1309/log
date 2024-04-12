package log

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type LogLevel int

var (
	DebugLevel LogLevel = 0
	InfoLevel  LogLevel = 1
	WarnLevel  LogLevel = 2
	ErrorLevel LogLevel = 3
)

var levelPrefix = []string{
	"[DEBUG] ",
	"[INFO ] ",
	"[WARN ] ",
	"[ERROR] ",
}

type Logger struct {
	level      LogLevel
	baseLogger *log.Logger
	baseFile   *os.File
}

func New(level LogLevel, pathname string, flag int) (*Logger, error) {
	var baseLogger *log.Logger
	var baseFile *os.File
	if pathname != "" {
		now := time.Now()
		filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		// if dir not exist, create new dir
		if _, err := os.Stat(pathname); err != nil {
			if os.IsNotExist(err) {
				if e := os.Mkdir(pathname, os.ModePerm); e != nil {
					return nil, e
				}
			}
		}

		file, err := os.Create(path.Join(pathname, filename))
		if err != nil {
			return nil, err
		}

		baseLogger = log.New(file, "", flag)
		baseFile = file
	} else {
		baseLogger = log.New(os.Stdout, "", flag)
	}

	logger := new(Logger)
	logger.level = level
	logger.baseLogger = baseLogger
	logger.baseFile = baseFile

	return logger, nil
}

func (l *Logger) Close() {
	if l.baseFile != nil {
		l.Close()
	}

	l.baseLogger = nil
	l.baseFile = nil
}

func (l *Logger) output(level LogLevel, format string, args ...any) {
	// skip
	if level < l.level {
		return
	}

	if l.baseLogger == nil {
		log.Fatal("log output err, baseLogger is nil")
		return
	}

	format = levelPrefix[level] + format
	l.baseLogger.Output(3, fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(format string, args ...any) {
	l.output(DebugLevel, format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.output(InfoLevel, format, args...)
}

func (l *Logger) Warn(format string, args ...any) {
	l.output(WarnLevel, format, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.output(ErrorLevel, format, args...)
}

var defaultLogger, _ = New(DebugLevel, "", log.LstdFlags)

func Debug(format string, args ...any) {
	defaultLogger.Debug(format, args...)
}

func Info(format string, args ...any) {
	defaultLogger.Info(format, args...)
}

func Warn(format string, args ...any) {
	defaultLogger.Warn(format, args...)
}

func Error(format string, args ...any) {
	defaultLogger.Error(format, args...)
}
