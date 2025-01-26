// code-buddy-consumer/main.go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
    conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my-topic", 0)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    for {
        message, err := conn.ReadMessage(1e3)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Received:", string(message.Value))
    }
}
