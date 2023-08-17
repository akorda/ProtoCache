package main

import (
	"log"
	"os"

	"github.com/akorda/protocache/caching"
	"github.com/akorda/protocache/server"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cache := caching.NewMemoryDistributedCache()

	options := server.ProtoCacheOptions{
		ListenAddress: ":4000",
	}
	server, err := server.NewProtoCacheServer(cache, options)
	if err != nil {
		log.Fatal(err)
	}

	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
