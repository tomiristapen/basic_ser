package natsinfra

import (
    "log"

    "github.com/nats-io/nats.go"
)

func ConnectNATS(url string) *nats.Conn {
    nc, err := nats.Connect(url)
    if err != nil {
        log.Fatalf("NATS connection failed: %v", err)
    }
    return nc
}