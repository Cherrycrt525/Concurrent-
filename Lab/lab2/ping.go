package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	for {
		fmt.Println("\nFoo is sending: ping")
		trace.Log(context.Background(), "Sent", "ping")
		channel <- "ping"

		message := <-channel
		fmt.Println("Foo has received:", message)
		trace.Log(context.Background(), "Received", message)
	}
}

func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
	for {
		message := <-channel
		fmt.Println("\nBar has received:", message)
		trace.Log(context.Background(), "Received", message)

		fmt.Println("Bar is sending: pong")
		trace.Log(context.Background(), "Sent", "Pong")
		channel <- "pong"
	}
}

func pingPong() {
	pingPong := make(chan string)
	// TODO: make channel of type string and pass it to foo and bar
	go foo(pingPong) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(pingPong)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
