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

func createLogDirectory() {
	// Define the logs directory
	logsDir := "logs"

	// Check if the logs directory exists
	if _, err := os.Stat(logsDir); os.IsExist(err) {
		return
	}

	// Create the logs directory if it doesn't exist
	err := os.Mkdir(logsDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating logs directory: %v\n", err)
		return
	}
}
