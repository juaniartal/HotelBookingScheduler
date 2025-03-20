package main

import (
	"net/http"
	"testing"
)

package backend_test

import (
    "testing"
    "net/http"
    "log"
    "os/exec"
    "time"
)

func TestServer(t *testing.T) {
    // Inicia el servidor en segundo plano
    cmd := exec.Command("go", "run", "main.go", "database.go", "handlers.go", "routes.go", "models.go")
    err := cmd.Start()
    if err != nil {
        t.Fatalf("Failed to start server: %v", err)
    }

    // Espera un momento para asegurarse de que el servidor esté listo
    time.Sleep(2 * time.Second) // Ajusta el tiempo según sea necesario

    // Realiza la solicitud GET
    resp, err := http.Get("http://localhost:8080/bookings")
    if err != nil {
        t.Fatalf("Failed to make GET request: %v", err)
    }
    defer resp.Body.Close()

    // Verifica la respuesta del servidor
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200 OK, got %v", resp.StatusCode)
    }

    // Detiene el servidor después de la prueba
    err = cmd.Process.Kill()
    if err != nil {
        t.Fatalf("Failed to kill server: %v", err)
    }
}
