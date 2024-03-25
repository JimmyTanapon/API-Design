package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serverMux() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`mux jow lai`))
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello ! World...`))
	})

	srv := http.Server{
		Addr:    ":2565",
		Handler: mux,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	fmt.Println("Server starting at :2565 ")
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	fmt.Println("Shutting down.... !")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("shutdown error!:", err)
	}
	fmt.Println("server shutdown ..........")
}
