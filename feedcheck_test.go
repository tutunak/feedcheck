package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock HTTP server for testing
func createMockServer(headerValue string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if headerValue != "" {
			w.Header().Add("Last-Modified", headerValue)
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func TestLastModified(t *testing.T) {
	// 1. Test a successful scenario
	server := createMockServer("Mon, 02 Jan 2023 15:04:05 GMT")
	defer server.Close()
	result, err := lastModified(server.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result != "Mon, 02 Jan 2023 15:04:05 GMT" {
		t.Errorf("Expected header 'Mon, 02 Jan 2023 15:04:05 GMT', but got: %v", result)
	}

	// 2. Test scenario where the header is missing
	serverMissingHeader := createMockServer("")
	defer serverMissingHeader.Close()
	result2, err2 := lastModified(serverMissingHeader.URL)
	if err2 != nil {
		t.Errorf("Expected no error, but got: %v", err2)
	}
	if result2 != "" {
		t.Errorf("Expected empty header, but got: %v", result2)
	}

	// 3. Test an erroneous scenario (e.g., invalid URL)
	_, err3 := lastModified("invalid_url")
	if err3 == nil {
		t.Error("Expected an error for invalid URL but got none")
	}
}
