package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Load the config.yaml file
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// Create an instance of SystemMetrics which implements the MetricsCollector interface
	var collector MetricsCollector = &SystemMetrics{}

	for {
		// Collect metrics using the interface methods
		cpuUsage := collector.GetCPUUsage()
		memoryUsage := collector.GetMemoryUsage()
		diskUsage := collector.GetDiskUsage()

		// Log the collected metrics, pass log directory
		logMetrics(cpuUsage, memoryUsage, diskUsage, config.LogDirectory)

		// Check if metrics exceed thresholds
		var notificationMessage string
		if cpuUsage > config.CPUUsageThreshold {
			notificationMessage += fmt.Sprintf("%v: ⚠️ CPU usage too high: %d%%\n", GetHostname(), cpuUsage)
		}
		if memoryUsage > config.MemoryUsageThreshold {
			notificationMessage += fmt.Sprintf("%v: ⚠️ Memory usage too high: %d%%\n", GetHostname(), memoryUsage)
		}
		if diskUsage > config.DiskUsageThreshold {
			notificationMessage += fmt.Sprintf("%v: ⚠️ Disk usage too high: %d%%\n", GetHostname(), diskUsage)
		}

		// Send LINE notification if needed
		if notificationMessage != "" {
			err := sendLineNotification(notificationMessage)
			if err != nil {
				// Log the error to the log file
				logError(err.Error(), config.LogDirectory) // Log the error to the log file
				log.Printf("Failed to send LINE notification: %v", err)
			} else {
				fmt.Println("Notification sent successfully.")
			}
		} else {
			fmt.Println("Metrics are within thresholds, no notification sent.")
		}

		// Sleep for the configured interval
		time.Sleep(time.Duration(config.IntervalSeconds) * time.Second)
	}
}
