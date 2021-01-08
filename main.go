package main

import (
	"fmt"
	"goproxy/services"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const appVersion = "3.0"

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("err : %s", err)
	}
	cleanUp(&service.S)
}

func cleanUp(s *services.Service) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for range signalChan {
			fmt.Println("\nReceived an interrupt, stopping services...")
			(*s).Clean()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
