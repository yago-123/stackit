package main

import (
	"log"

	"github.com/yago-123/stackit/config"
	v1Api "github.com/yago-123/stackit/pkg/api/v1"
	"github.com/yago-123/stackit/pkg/encoding"
	"github.com/yago-123/stackit/pkg/storage"
)

func main() {
	// todo(): config should be provided via argument or env
	// todo(): something like $ go run main.go --config default-config.yml
	cfg := config.New()

	// todo(): detach json encoder from storage if possible
	jsonEncoder := encoding.NewJson()

	store := storage.NewMemory(cfg, jsonEncoder)

	// todo(): spawn in a different go routine and use wait
	server := v1Api.New(cfg, store)
	if err := server.Start(); err != nil {
		// todo(): use cfg.Logger instead
		log.Fatalf("failed to start server: %v", err)
	}

	defer server.Stop()
}
