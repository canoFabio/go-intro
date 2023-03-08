package fallbacks

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const maxErrors = 3
const openTimeout = 10 * time.Second

var state = "closed"
var errorCount = 0
var mutex = &sync.Mutex{}

func publishLogs(logs []string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if state == "abierto" {
		return fmt.Errorf("Circuit Breaker open")
	}

	err := postLogs(logs)
	if err != nil {
		errorCount++
		if errorCount >= maxErrors {
			tripCircuitBreaker()
		}
		saveLogsToFile(logs)
		return fmt.Errorf("Falló la publicación de los logs en el servidor remoto")
	}

	resetCircuitBreaker()
	return nil
}

func postLogs(logs []string) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post("http://logs.example.com", "text/plain", bytes.NewBufferString(strings.Join(logs, "\n")))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func tripCircuitBreaker() {
	state = "fail"
	time.AfterFunc(openTimeout, func() {
		state = "open"
		errorCount = 0
	})
}

func saveLogsToFile(logs []string) error {
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo para guardar los logs: %s", err.Error())
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(strings.Join(logs, "\n")); err != nil {
		log.Printf("Error al guardar los logs en el archivo: %s", err.Error())
		return err
	}
	return nil
}

func resetCircuitBreaker() {
	mutex.Lock()
	defer mutex.Unlock()
	state = "closed"
	errorCount = 0
}
