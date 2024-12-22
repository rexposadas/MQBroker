package main

import (
	"log"

	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

func main() {
	// Create a new MQTT server instance
	server := mqtt.New()

	port := ":1884"

	// Create a TCP listener on port 1883
	tcp := listeners.NewTCP("tcp1", port)

	// Add the listener to the server
	err := server.AddListener(tcp, nil)
	if err != nil {
		log.Fatalf("Failed to add TCP listener: %v", err)
	}

	// Start the server
	go func() {
		log.Printf("Starting MQTT broker on port %s...\n", port)
		if err := server.Serve(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for a signal to exit (e.g., Ctrl+C)
	select {}
}
