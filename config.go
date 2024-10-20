package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config structure to hold thresholds and interval
type Config struct {
	CPUUsageThreshold    int    `yaml:"cpu_usage_threshold"`
	MemoryUsageThreshold int    `yaml:"memory_usage_threshold"`
	DiskUsageThreshold   int    `yaml:"disk_usage_threshold"`
	IntervalSeconds      int    `yaml:"interval_seconds"` // New field for the interval
	LogDirectory         string `yaml:"log_directory"`    // New field for log directory
}

// LoadConfig loads the thresholds and interval from the config.yaml file
func LoadConfig() (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile("config.yaml") // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	// Default to 5 seconds if interval_seconds is not specified
	if config.IntervalSeconds == 0 {
		config.IntervalSeconds = 5
	}

	return config, nil
}
