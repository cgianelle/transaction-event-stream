package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "transactions"
	brokerAddress = "localhost:9092"
)

func main() {
	fmt.Println("Transaction Generator v1.0")
	ctx := context.Background()

	generateTransactions(ctx)
}

func generateTransactions(ctx context.Context) {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	for {
		// Generate the random transaction value between -100 to 100
		n := -100 + rand.Intn(100 - -100 + 1)
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte("A3545FWF343 0001"),
			// create an arbitrary message payload for the value
			Value: []byte(strconv.Itoa(n)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// sleep for a second
		time.Sleep(time.Second)
	}

}
