package utils

import (
	"fmt"
	"os"
	"time"
)

// WriteLog escribe un mensaje al archivo de log con timestamp.
func WriteLog(message string) error {
	file, err := os.OpenFile("dev-debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] %s\n", timestamp, message)

	_, err = file.WriteString(logMessage)
	return err
}

// LogError registra un error en el archivo de log.
func LogError(err error) {
	if err != nil {
		WriteLog("ERROR: " + err.Error())
	}
}

// LogInfo registra información general en el archivo de log.
func LogInfo(message string) {
	WriteLog("INFO: " + message)
}

// LogDebug registra información de debug en el archivo de log.
func LogDebug(message string) {
	WriteLog("DEBUG: " + message)
}

// LogRequest registra detalles de una petición HTTP.
func LogRequest(method, url string, statusCode int) {
	message := fmt.Sprintf("REQUEST: %s %s - Status: %d", method, url, statusCode)
	WriteLog(message)
}
