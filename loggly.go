// Package loggly provides an interface for sending messages to Loggly using
// the users credentials supplied by an environment variable
package loggly

// Import packages
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	color "github.com/TwiN/go-color"
)

// Loggly API URL
var logglyURL = "http://logs-01.loggly.com/inputs/"

// Permitted levels
var levelsAllowed = []string{"error", "warn", "info", "verbose", "debug", "silly"}

// ClientType stores the URL and tag
type ClientType struct {
	URL string
	Tag string
}

// messageType conforms to the JSON expected by Loggly
type messageType struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// New instantiates a new Loggly client using the
// tag provided by the caller
func New(tag string) *ClientType {

	// Check for LOGGLY_TOKEN enviroment variable
	token := os.Getenv("LOGGLY_TOKEN")
	if token == "" {
		// Variable undefined
		fmt.Print(color.Ize(color.Yellow, "WARNING: "))
		fmt.Println("LOGGLY_TOKEN environment variable undefined. No messages can be sent to Loggly")
	}
	newClient := &ClientType{}
	newClient.URL = logglyURL + token + "/tag/" + tag
	newClient.Tag = tag

	return newClient
}

// Send constructs and sends a message to loggly.
func (c *ClientType) Send(level string, message string) error {

	// Validate the level
	if !checkLevel(level) {
		return errors.New("Invalid level: " + level)
	}

	// build the message
	var msg messageType
	msg.Level = level
	msg.Message = message

	// Timestamp the message
	msg.Timestamp = time.Now().Truncate(time.Millisecond).UTC().String()

	// Send message to loggly
	c.sendToLoggly(msg)

	return nil
}

// EchoSend sends a message to loggly and logs the
// message to the console
func (c *ClientType) EchoSend(level string, message string) error {

	err := c.Send(level, message)
	if err != nil {
		return err
	}

	// log the message to the console
	log.Println(level, message)

	return nil
}

// non-exported functions

// sendToLoggly sends message object contents to Loggly
func (c *ClientType) sendToLoggly(msg messageType) error {

	// Marshal message object to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return errors.New("Could not create JSON payload")
	}

	// Construct the HTTP request with timeout
	httpClient := &http.Client{
		Timeout: 7 * time.Second}
	request, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(msgBytes))
	if err != nil {
		return err
	}

	// Set the header attributes
	request.Header.Add("Content-Type", "application/json")

	_, err = httpClient.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// This function validates a message level
func checkLevel(level string) bool {

	for _, testLevel := range levelsAllowed {
		if testLevel == level {
			// Match -- OK
			return true
		}
	}

	return false
}
