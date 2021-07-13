package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "transactions"
	brokerAddress = "localhost:9092"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	consume(ctx)
}

func consume(ctx context.Context) {
	var ledger int = 0
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerAddress},
		Topic:       topic,
		StartOffset: 0,
		Partition:   0,
		// assign the logger to the reader
		Logger: l,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		value := string(msg.Value)
		transaction, _ := strconv.Atoi(value)
		fmt.Printf("offset: %v, received: %v for %v\n", msg.Offset, transaction, string(msg.Key))
		ledger = ledger + transaction
		fmt.Printf("Ledger Balance: %v\n", ledger)
	}

}
