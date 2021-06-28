package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"nuorder/handler"
)

func main() {
	h := handler.Handler{}
	if err := h.New(); err != nil {
		log.Fatalf("error initializing handler, ERROR: %s\n", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":8443", h.Router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server, ERROR: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")
}
