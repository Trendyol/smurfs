package logger

import "fmt"

type DefaultLogger struct {
}

func (d DefaultLogger) Debug(message string, args ...interface{}) {
	fmt.Printf("DHOST-DEBUG: %s\n", message)
}

func (d DefaultLogger) Info(message string, args ...interface{}) {
	fmt.Printf("DHOST-INFO: %s\n", message)
}

func (d DefaultLogger) Warn(message string, args ...interface{}) {
	fmt.Printf("DHOST-WARN: %s\n", message)
}

func (d DefaultLogger) Error(message string, args ...interface{}) {
	fmt.Printf("DHOST-ERROR: %s\n", message)
}

func (d DefaultLogger) Fatal(message string, args ...interface{}) {
	fmt.Printf("DHOST-FATAL: %s\n", message)
}
