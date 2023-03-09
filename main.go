package main

import (
	"context"
	"fmt"
	"golang_test_task/server"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	fmt.Println("Initializing .env and other things...")
}

func main() {
	fmt.Println("Starting HTTP server...")
	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	done := make(chan struct{})
	defer close(done)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		server.InitServer(ctx)
		done <- struct{}{}
	}()

	<-sigs
	cancel()
	<-done
}
