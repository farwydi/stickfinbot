package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	endpoint, cleanup, err := bootstrap()
	if err != nil {
		log.Fatalf("Fail bootstrap app: %v", err)
	}
	defer cleanup()

	go func() {
		if err := endpoint.Run(); err != nil {
			log.Fatalf("Fail run app: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	endpoint.Stop()
	log.Println("Power off")
}
