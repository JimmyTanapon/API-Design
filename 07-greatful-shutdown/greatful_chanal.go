package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func signalFucn() {
	log.Println("Server Start..!")

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM) //os.interupt รอดักinteruptที่เกิดขึ้น
	log.Println("wait for signal!")
	<-stop

	log.Println("Server Stopped!!!!")

}
