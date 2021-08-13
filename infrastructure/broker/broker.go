package broker

import (
	"context"

	"github.com/nats-io/nats.go"
)

var (
	messageBroker *MessageBroker
)

// MessageBroker the message broker
type MessageBroker struct {
	*nats.Conn
	URI string
}

// NewClient prepare message broker
func NewClient(uri string) *MessageBroker {

	if messageBroker != nil {

		return messageBroker
	}

	messageBroker = &MessageBroker{
		URI: uri,
	}
	return messageBroker
}

// GetClient gets message broker
func GetClient() *MessageBroker {

	if messageBroker != nil {
		return messageBroker
	}
	return nil
}

// Connect connect to message broker
func (m *MessageBroker) Connect() error {
	var err error

	m.Conn, err = nats.Connect(m.URI)

	return err
}

// Flush flush all messages in queue
func (m *MessageBroker) Flush(ctx context.Context) error {
	return m.FlushWithContext(ctx)
}
