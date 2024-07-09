package configs

import (
	"fmt"
	"log"
	"os"
	"time"
)

func InitLogger() {
	createLogDirectory()
	// Get the current date
	currentTime := time.Now()

	// Format the date
	logFileName := fmt.Sprintf("logs/log_%02d-%s-%d.txt", currentTime.Day(), currentTime.Month().String()[:3], currentTime.Year())

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

// createLogDirectory ensures that the "logs" directory exists.
// If the directory does not exist, it creates it with appropriate permissions.
func createLogDirectory() {
	// Define the logs directory
	logsDir := "logs"
	exist, _ := os.Stat(logsDir)
	if exist != nil {
		return
	}

	// Create the logs directory if it doesn't exist
	err := os.Mkdir(logsDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating logs directory: %v\n", err)
	}
}
