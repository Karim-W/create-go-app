package rabbit

import (
	"os"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	rmqURI := os.Getenv("RMQ_URI")
	if rmqURI == "" {
		t.Skip("RMQ_URI not set")
	}
	conn, channel, err := MustSetupRMQConnection("test", rmqURI)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(30 * time.Second)
	channel.Close()
	conn.Close()
}
