package handlers

import (
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../../frontend/templates/index.html")
}

func AgentHandler(w http.ResponseWriter, r *http.Request) {
    // Placeholder for agent-related operations
    w.Write([]byte("Agent endpoint"))
}