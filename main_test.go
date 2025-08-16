package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// TestRootHandler tests the root endpoint
func TestRootHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	newrelicApp, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("go-new-relic-test"),
		newrelic.ConfigLicense("0123456789012345678901234567890123456789"),
	)
	logger := log.New(log.Writer(), "", log.Default().Flags())
	server := server{newrelicApp: newrelicApp, logger: logger}

	// Act
	server.rootHandler(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: expected %v, received %v",
			http.StatusOK, status)
	}
	expected := "The server is running"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned a wrong body: expected %v, received %v",
			expected, rr.Body.String())
	}
}

// TestReportHandler tests the /report
func TestReportHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest("GET", "/report", nil)
	rr := httptest.NewRecorder()
	newrelicApp, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("go-new-relic-test"),
		newrelic.ConfigLicense("0123456789012345678901234567890123456789"),
	)
	logger := log.New(log.Writer(), "", log.Default().Flags())
	server := server{newrelicApp: newrelicApp, logger: logger}

	// Act
	server.reportHandler(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: expected %v, received %v",
			http.StatusOK, status)
	}
	expected := "Report generated successfully!"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned a wrong body: expected %v, received %v",
			expected, rr.Body.String())
	}
}
