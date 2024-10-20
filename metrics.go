package main

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// MetricsCollector interface defines methods for collecting system metrics
type MetricsCollector interface {
	GetCPUUsage() int
	GetMemoryUsage() int
	GetDiskUsage() int
}

// SystemMetrics implements the MetricsCollector interface
type SystemMetrics struct{}

// GetCPUUsage retrieves the CPU usage percentage
func (s *SystemMetrics) GetCPUUsage() int {
	var cpuUsage float64

	// Switch case to handle different operating systems
	switch runtime.GOOS {
	case "windows":
		// Windows-specific logic (if needed in the future)
		cpuUsages, err := cpu.Percent(0, false)
		if err != nil {
			log.Fatalf("Error getting CPU usage: %v", err)
		}
		cpuUsage = cpuUsages[0]
	case "darwin": // macOS
		// macOS-specific logic (if needed in the future)
		cpuUsages, err := cpu.Percent(0, false)
		if err != nil {
			log.Fatalf("Error getting CPU usage: %v", err)
		}
		cpuUsage = cpuUsages[0]
	default:
		log.Fatalf("Unsupported OS: %s", runtime.GOOS)
	}

	return int(cpuUsage)
}

// GetMemoryUsage retrieves the memory usage percentage
func (s *SystemMetrics) GetMemoryUsage() int {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error getting memory usage: %v", err)
	}

	// Switch case to handle different operating systems
	switch runtime.GOOS {
	case "windows":
		// Windows-specific logic (if needed in the future)
		return int(vmStat.UsedPercent)
	case "darwin": // macOS
		// macOS-specific logic (if needed in the future)
		return int(vmStat.UsedPercent)
	default:
		log.Fatalf("Unsupported OS: %s", runtime.GOOS)
		return -1 // Fallback return value for unsupported OS
	}
}

// GetDiskUsage retrieves the disk usage percentage
func (s *SystemMetrics) GetDiskUsage() int {
	var cmd *exec.Cmd
	var out []byte
	var err error

	// Switch case to handle different operating systems
	switch runtime.GOOS {
	case "windows":
		// Windows-specific logic
		cmd = exec.Command("wmic", "logicaldisk", "get", "size,freespace,caption")
		out, err = cmd.Output()
		if err != nil {
			log.Fatalf("Error getting disk usage: %v", err)
		}
		// Parse output to calculate used percentage (simple example)
		// This may require more parsing based on actual output format
		lines := strings.Split(string(out), "\n")
		if len(lines) > 1 {
			parts := strings.Fields(lines[1]) // First line after header
			if len(parts) == 3 {              // Size and Free Space are available
				totalSpace, _ := strconv.ParseInt(parts[1], 10, 64)
				freeSpace, _ := strconv.ParseInt(parts[2], 10, 64)
				usedSpace := totalSpace - freeSpace
				return int((float64(usedSpace) / float64(totalSpace)) * 100)
			}
		}
	case "darwin": // macOS
		// macOS-specific logic
		cmd = exec.Command("df", "-H")
		out, err = cmd.Output()
		if err != nil {
			log.Fatalf("Error getting disk usage: %v", err)
		}
		// Parse output to get used percentage
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "/") { // Assuming "/" is the root directory
				fields := strings.Fields(line)
				if len(fields) > 4 { // Ensure the expected fields are available
					usage := fields[4] // The usage percentage is usually in the 5th column
					usagePercent := strings.TrimSuffix(usage, "%")
					return atoi(usagePercent) // Convert to int
				}
			}
		}
	default:
		log.Fatalf("Unsupported OS: %s", runtime.GOOS)
		return -1 // Fallback return value for unsupported OS
	}

	return -1 // Fallback if parsing fails
}

// Helper function to convert string to int
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
