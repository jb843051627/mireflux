package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jb843051627/mireflux/internal/api"
	"github.com/jb843051627/mireflux/internal/service"
	"github.com/jb843051627/mireflux/internal/store"
)

func main() {
	databasePath := os.Getenv("MIREFLUX_PROBE_DB")
	if databasePath == "" {
		log.Fatal("MIREFLUX_PROBE_DB is required")
	}
	repository, err := store.Open(databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer repository.Close()
	app := service.NewLab(repository)
	defer app.Close()
	server := httptest.NewServer(api.New(app))
	defer server.Close()

	response, err := http.Get(server.URL + "/healthz")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("health status: %s", response.Status)
	}
	var health struct {
		Database string `json:"database"`
	}
	if err := json.NewDecoder(response.Body).Decode(&health); err != nil {
		log.Fatal(err)
	}
	if health.Database != databasePath {
		log.Fatalf("health database %q does not match requested path %q", health.Database, databasePath)
	}
	if _, err := os.Stat(databasePath); err != nil {
		log.Fatal(err)
	}
	if err := repository.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mireflux probe passed: %s\n", databasePath)
}
