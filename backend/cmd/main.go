package main

import (
    "log"
    "net/http"

    "radio/backend/internal/handlers"
    "radio/backend/internal/utils"
)

func main() {
    // Load configuration
    config, err := utils.LoadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    // Initialize routes
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/api/agent", handlers.AgentHandler)

    // Start server
    log.Printf("Server starting on %s", config.ServerAddress)
    if err := http.ListenAndServe(config.ServerAddress, nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}