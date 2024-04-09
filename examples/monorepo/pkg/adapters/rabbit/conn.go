package rabbit

import (
	"log"
	"net"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MustSetupRMQConnection(
	name string,
	amqpURI string,
) (*amqp.Connection, *amqp.Channel, error) {
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName(name)

	log.Printf("dialing %q", amqpURI)

	conn, err := amqp.DialConfig(amqpURI, config)
	if err != nil {
		log.Fatalf("Dial: %s", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Channel: %s", err)
	}

	go func() {
		pingTicker := time.NewTicker(5 * time.Second)
		defer pingTicker.Stop()
		for {
			<-pingTicker.C
			// remote address
			remote := conn.RemoteAddr().String()
			_, err := net.DialTimeout("tcp", remote, 3*time.Second)
			if err != nil {
				log.Fatalf("heartbeat: %s", err.Error())
			}
		}
	}()

	return conn, channel, nil
}
