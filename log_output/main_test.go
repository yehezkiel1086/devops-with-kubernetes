package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatusEndpoint(t *testing.T) {
	testUUID := "test-uuid-1234"
	router := newRouter(testUUID)

	tests := []struct {
		name string
		path string
	}{
		{name: "root path", path: "/"},
		{name: "status path", path: "/status"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
			}

			body := rec.Body.String()
			if !strings.Contains(body, testUUID) {
				t.Errorf("expected body to contain '%s', got: %s", testUUID, body)
			}

			contentType := rec.Header().Get("Content-Type")
			if !strings.Contains(contentType, "text/plain") {
				t.Errorf("expected text/plain content type, got: %s", contentType)
			}
		})
	}
}
