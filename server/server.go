package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	PORT string = ":5000"
)

func main() {
	clientPath := filepath.Join("client")
	fileServer := http.FileServer(http.Dir(clientPath))

	http.Handle("/", fileServer)

	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON body
		var msg Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		// Append data to a file
		filepath := filepath.Join("..", "logs", "msgs.txt")
		if err := appendToFile(filepath, msg.ContactInfo, msg.Message); err != nil {
			http.Error(w, "Failed to save message", http.StatusInternalServerError)
			log.Printf("Error writing to file: %v", err)
			return
		}

		log.Printf("Received data: %+v", msg)

		w.WriteHeader(http.StatusOK)
	})

	log.Printf("Server started on port %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Message struct {
	ContactInfo string `json:"contact_info"`
	Message     string `json:"message"`
}

func appendToFile(filename, contact, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Format the entry
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// entry := "[" + timestamp + "] Email: " + contact + "\nMessage: " + message + "\n\n"
	entry := fmt.Sprintf("%s\n\nEmail:\n%s\n\nMessage:\n%s\n", timestamp, contact, message)
	entry += "-----------------\n"

	// Write the entry to the file
	_, err = file.WriteString(entry)
	return err
}
