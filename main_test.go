package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
	"log"
)

var (
	mockUrls = []string{"https://www.google.com", "https://www.example.com"}
)

func TestMain(m *testing.M) {
	// Set the environment variable for testing
	os.Setenv("PING_URLS", strings.Join(mockUrls, ","))
	defer os.Unsetenv("PING_URLS") // Clean up after test

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestPingURL(t *testing.T) {
	for _, url := range mockUrls {
		// Create a mock server for each URL
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		// Update URL with the mock server address
		url = strings.Replace(url, "https://", ts.URL+"/", -1)

		// Capture log output
		var logBuffer bytes.Buffer
		log.SetOutput(&logBuffer)
		defer func() {
			log.SetOutput(os.Stdout)
		}()

		// Run the ping function in a separate goroutine
		go pingURL(url)

		// Wait for some time to simulate pinging
		time.Sleep(time.Second * 2)

		// Check for expected logs
		expected := "Pinging " + url
		if !strings.Contains(logBuffer.String(), expected) {
			t.Errorf("Expected log for pinging %s not found", url)
		}

		// Check for error logs (assuming successful response)
		if strings.Contains(logBuffer.String(), "Error:") {
			t.Errorf("Unexpected error log for %s", url)
		}
	}
}

func pingURL(url string) {
	url = strings.TrimSpace(url)
	for {
		_, err := http.Get(url)
		log.Println("Pinging " + url)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		time.Sleep(time.Second * 5)
	}
}