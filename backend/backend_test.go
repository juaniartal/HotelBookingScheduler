package main

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// Aquí asumes que el servidor se ejecuta en localhost:8080
	resp, err := http.Get("http://localhost:8080") // Cambia esta URL si tu servidor está en otro puerto
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}
