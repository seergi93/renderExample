package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Puerto para Render (PORT env var, fallback 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] Health check desde %s\n", time.Now().Format(time.RFC3339), r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	// Root endpoint que solo loguea
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logEntry := fmt.Sprintf("[%s] Petición recibida: %s %s desde %s - User-Agent: %s",
			time.Now().Format(time.RFC3339),
			r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		fmt.Println(logEntry) // Imprime por consola (stdout)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Servicio activo - logs en consola")
	})

	fmt.Printf("Servicio iniciado en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
