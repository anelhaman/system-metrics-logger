# System Metrics Logger with LINE Notify
This Golang script collects system metrics from various platforms (Windows, Mac Intel, Mac M1, Linux Ubuntu, CentOS) and logs them into a file. It also integrates with LINE Notify to send alerts when specific thresholds (set in a YAML file) are exceeded. The LINE Notify token is securely stored in a `.env` file.

## Features
* Collects system metrics (CPU, memory, disk usage, etc.).
* Logs metrics into a designated file.
* Sends notifications via LINE Notify when metrics exceed set thresholds.
* Thresholds for each metric are customizable in a `config.yaml` file.
* Supports the following platforms:
    * Windows
    * macOS (Intel and M1 architectures)
    * Linux (Ubuntu, CentOS)
* Secure storage of sensitive information (LINE Notify token) in a `.env` file.

## Installation
1. Clone the repository:
```
git clone https://github.com/anelhaman/system-metrics-logger.git
cd system-metrics-logger
```
2. Install Golang dependencies using Go modules:

```
go mod tidy
```
3. Create a `.env` file in the root directory and add your LINE Notify token:

```
LINE_NOTIFY_TOKEN=your-line-notify-token
```
4. Create a config.yaml file for the threshold values:
```
cpu_usage_threshold: 80    # Threshold for CPU usage (in percentage)
memory_usage_threshold: 70 # Threshold for memory usage (in percentage)
disk_usage_threshold: 90   # Threshold for disk usage (in percentage)
log_directory: ""          # Log directory (default is current directory)
```

## Usage
Run the script using:

```
go run main.go
```
The script will log system metrics into `metrics.log` and send LINE notifications if thresholds are exceeded.

## Example of running on different systems:
* Windows: Run the above command in a PowerShell or CMD.
* macOS: Run in Terminal.
* Linux: Run in any terminal or bash shell.

# Environment Variables
The following environment variables need to be set in the .env file:


| Variable | Description|
| ------------- |:-------------:|
| LINE_NOTIFY_TOKEN | Your LINE Notify token for sending alerts.


## Configuration
The script uses a config.yaml file to specify metric thresholds. You can adjust these thresholds according to your system's needs.

Example config.yaml:

```
cpu_usage_threshold: 80    # in percentage
memory_usage_threshold: 75 # in percentage
disk_usage_threshold: 90   # in percentage
log_directory: ""          # Log directory (default is current directory)
```

Create a `.env` file in the root directory of the project using the provided `.env.example` as a template. 

### Example `.env` file:

```
# .env

# LINE Notify token for sending notifications
LINE_NOTIFY_TOKEN=your_line_notify_token_here

```

## Platform Support
* Windows: All versions.
* macOS: Both Intel and Apple Silicon (M1) architectures.
* Linux:
    * Ubuntu
    * CentOS

## LINE Notify Integration
To enable LINE notifications, you need to generate a LINE Notify token:

1. Visit LINE Notify.
2. Log in and generate a personal access token.
3. Store this token in the .env file as described above.
Whenever a metric exceeds its threshold, a notification will be sent to your LINE app with details about the exceeded metric.

## Logging
All system metrics are logged into metrics.log in the root directory. The log contains information like:

* Timestamp of the log entry.
* CPU usage.
* Memory usage.
* Disk usage.

## Example log entry:

```
2024-01-01 12:00:00 | CPU: 60% | Memory: 65% | Disk: 55%
``` 

## Building the Application
You can build the executable for your system using the following commands:

### Required Programs for Native Go Build
#### Go Programming Language:

* Download and Install Go:
    * For Windows: Download the Windows installer from the [official Go website](https://go.dev/doc/install). Run the installer and follow the prompts to complete the installation.
    * For macOS: Download the macOS installer from the [official Go website](https://go.dev/doc/install). You can also install Go using Homebrew
    ```
    $ brew install go
    ```

### Build Command
#### For Windows:

```
GOOS=windows GOARCH=amd64 go build -o system-metrics-logger.exe
```
#### For macOS (Intel):

```
GOOS=darwin GOARCH=amd64 go build -o system-metrics-logger
```
#### For macOS (M1):

```
GOOS=darwin GOARCH=arm64 go build -o system-metrics-logger
```
#### For Linux (Ubuntu):

```
GOOS=linux GOARCH=amd64 go build -o system-metrics-logger
```
#### For Linux (CentOS):

```
GOOS=linux GOARCH=amd64 go build -o system-metrics-logger
```
## License
This project is licensed under the MIT License.