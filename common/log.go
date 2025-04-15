package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"sync"
)

// Config holds options for logger setup

var (
	once       sync.Once
	logFile    *os.File
	defaultLog *logrus.Logger
	logFileMu  sync.Mutex
)

// Config holds options for logger setup
type Config struct {
	EnableFile  bool
	LogFilePath string
	UseJSON     bool
	LogLevel    logrus.Level
}

// InitLogger initializes a global logger
func InitLogger(cfg Config) {
	once.Do(func() {
		logger := logrus.New()

		var writers []io.Writer
		writers = append(writers, os.Stdout) // Always write to stdout

		if cfg.EnableFile && cfg.LogFilePath != "" {
			var err error
			logFileMu.Lock()
			logFile, err = os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			logFileMu.Unlock()
			if err == nil {
				writers = append(writers, logFile)
			} else {
				log.Printf("Failed to log to file: %v", err)
			}
		}

		logger.SetOutput(io.MultiWriter(writers...))

		if cfg.UseJSON {
			logger.SetFormatter(&logrus.JSONFormatter{})
		} else {
			logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		}

		logger.SetLevel(cfg.LogLevel)
		logger.AddHook(&ErrorToStderrHook{})

		defaultLog = logger
	})
}

// ErrorToStderrHook sends error and higher logs to stderr
type ErrorToStderrHook struct{}

func (h *ErrorToStderrHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

func (h *ErrorToStderrHook) Fire(entry *logrus.Entry) error {
	msg, err := entry.String()
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(os.Stderr, msg)
	return err
}

// Log returns the global logger instance
func Log() *logrus.Logger {
	if defaultLog == nil {
		InitLogger(Config{})
	}
	return defaultLog
}

// CloseLogFile closes the log file if it's open
func CloseLogFile() {
	logFileMu.Lock()
	defer logFileMu.Unlock()
	if logFile != nil {
		_ = logFile.Close()
		logFile = nil
	}
}
