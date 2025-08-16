package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type server struct {
	nrApp  *newrelic.Application
	logger *log.Logger
}

func InitServer() server {
	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-new-relic"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}

	if err := nrApp.WaitForConnection(5 * time.Second); err != nil {
		panic(err)
	}

	writer := logWriter.New(os.Stdout, nrApp)
	logger := log.New(&writer, "", log.Default().Flags())

	logger.Println("New Relic agent initialized successfully")

	return server{nrApp: nrApp, logger: logger}
}

func (s *server) respond(w http.ResponseWriter, response string) {
	s.logger.Println("HTTP Response: ", response)
	fmt.Fprint(w, response)
}

func (s *server) rootHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("The root route has been reached!")
	s.respond(w, "The server is running")
}

func (s *server) reportHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("Generating a report log and sending it to New Relic APM with manual instrumentation...")

	// Simulate a blocking operation that takes 1 second
	time.Sleep(1 * time.Second)

	s.respond(w, "Report generated successfully!")
}

func main() {
	server := InitServer()

	http.HandleFunc(newrelic.WrapHandleFunc(server.nrApp, "/", server.rootHandler))
	http.HandleFunc(newrelic.WrapHandleFunc(server.nrApp, "/report", server.reportHandler))

	server.logger.Printf("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
