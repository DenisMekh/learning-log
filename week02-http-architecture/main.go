package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	application, err := New()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		shutDownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := application.Stop(shutDownCtx); err != nil {
			log.Fatal(err)
		}
	}()
	signalCtx, signalCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer signalCancel()
	go func() {
		if err := application.Start(); err != nil {
			log.Fatal(err)
		}

	}()
	<-signalCtx.Done()
}
