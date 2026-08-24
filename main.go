package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jb843051627/mireflux/internal/api"
	"github.com/jb843051627/mireflux/internal/service"
	"github.com/jb843051627/mireflux/internal/store"
)

func main() {
	databasePath := os.Getenv("MIREFLUX_DB")
	if databasePath == "" {
		databasePath = "data/mireflux.db"
	}
	repository, err := store.Open(databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer repository.Close()

	app := service.NewLab(repository)
	defer app.Close()

	address := os.Getenv("MIREFLUX_ADDR")
	if address == "" {
		address = ":8080"
	}
	server := &http.Server{
		Addr:              address,
		Handler:           api.New(app),
		ReadHeaderTimeout: 5 * time.Second,
	}
	stopped := make(chan os.Signal, 1)
	signal.Notify(stopped, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("mireflux listening on %s", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server stopped: %v", err)
		}
	}()
	<-stopped
	shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdown); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
