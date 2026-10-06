package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetHome(t *testing.T) {
	router := NewRouter()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.Mux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("expected Content-Type text/html, got %s", contentType)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "Todo App") {
		t.Errorf("expected body to contain 'Todo App', got: %s", body)
	}
}
