package main

import (
	"net/http"
	"os/exec"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "database.go", "handlers.go", "routes.go", "models.go")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	time.Sleep(10 * time.Second)

	resp, err := http.Get("http://localhost:8080/bookings")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", resp.StatusCode)
	}

	err = cmd.Process.Kill()
	if err != nil {
		t.Fatalf("Failed to kill server: %v", err)
	}
}
