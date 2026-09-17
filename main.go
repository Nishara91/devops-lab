package main

import (
	"fmt"
	"net/http"
)

// HealthMessage returns the system status string (used for unit testing)
func HealthMessage() string {
	return "System is healthy and operational!"
}

func main() {
	// Serve static files (HTML/CSS) from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Health check endpoint for CI/CD and container monitoring
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "%s"}`, HealthMessage())
	})

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}