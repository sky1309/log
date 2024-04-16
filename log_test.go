package log

import (
	"log"
	"testing"
)

func TestDefaultLog(t *testing.T) {
	Info("are you ok")
	Info("today is %d", 6)
	Warn("keep it moving")
	Error("error msg")
}

func TestNewLog(t *testing.T) {
	logger, _ := New(InfoLevel, "logs/", log.LstdFlags)
	logger.Debug("foo debug") // skip
	logger.Info("foo info")
	logger.Warn("foo war")
	logger.Error("foo error")
}
