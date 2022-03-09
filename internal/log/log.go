package log

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	levelVarName    = "GKE_POLICY_LOG"
	pathVarName     = "GKE_POLICY_LOG_PATH"
	formatVarName   = "GKE_POLICY_LOG_FORMAT"
	defaultLogLevel = logrus.InfoLevel
)

var log = NewLogger()

type envProvider interface {
	Getenv(key string) string
}

type OsEnvProvider struct{}

func (p OsEnvProvider) Getenv(key string) string {
	return os.Getenv(key)
}

func NewLogger() *logrus.Logger {
	envProvider := OsEnvProvider{}
	logger := logrus.New()
	logger.SetLevel(getLogLevel(envProvider))
	return logger
}

func getLogLevel(p envProvider) logrus.Level {
	switch value := p.Getenv(levelVarName); strings.ToLower(value) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return defaultLogLevel
	}
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Print(args ...interface{}) {
	log.Print(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warning(args ...interface{}) {
	log.Warning(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}
