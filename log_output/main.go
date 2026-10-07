package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func newRouter(storedString string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	statusHandler := func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		response := fmt.Sprintf("%s: %s\n", timestamp, storedString)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(response))
	}

	r.Get("/", statusHandler)
	r.Get("/status", statusHandler)

	return r
}

func main() {
	// generate random UUID and store to memory
	storedString := uuid.New().String()
	fmt.Printf("Initialized with random string: %s\n\n", storedString)

	// output the stored random UUID every 5 seconds in background
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for t := range ticker.C {
			timestamp := t.UTC().Format("2006-01-02T15:04:05.000Z")
			fmt.Printf("%s: %s\n", timestamp, storedString)
		}
	}()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	r := newRouter(storedString)

	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
