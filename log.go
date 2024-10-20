package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// GetHostname returns the lowercase hostname of the machine
func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(hostname)
}

// Log the collected metrics
func logMetrics(cpu, mem, disk int, logDir string) {
	hostname := GetHostname()
	logFileName := fmt.Sprintf("%s-%s.log", hostname, time.Now().Format("20060102"))
	logFilePath := logDir

	if logDir == "" {
		// Default to current directory if not specified
		logFilePath = "."
	}

	// Create the full path for the log file
	fullLogFilePath := fmt.Sprintf("%s/%s", logFilePath, logFileName)

	f, err := os.OpenFile(fullLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logEntry := fmt.Sprintf("%s | CPU: %d%% | Memory: %d%% | Disk: %d%%\n",
		time.Now().Format("2006-01-02 15:04:05"), cpu, mem, disk)
	if _, err := f.WriteString(logEntry); err != nil {
		log.Fatal(err)
	}
}

// LogError logs error messages to the log file
func logError(message string, logDir string) {
	hostname := GetHostname()
	logFileName := fmt.Sprintf("%s-%s.log", hostname, time.Now().Format("20060102"))
	logFilePath := logDir

	if logDir == "" {
		logFilePath = "."
	}

	fullLogFilePath := fmt.Sprintf("%s/%s", logFilePath, logFileName)

	f, err := os.OpenFile(fullLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logEntry := fmt.Sprintf("%s | ERROR: %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
	if _, err := f.WriteString(logEntry); err != nil {
		log.Fatal(err)
	}
}
