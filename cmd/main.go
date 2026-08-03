package main

import (
	"context"
	"ship/internal/client"
	"ship/internal/config"
	"ship/internal/server"
)

func main() {
	cfg, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := client.New(ctx, cfg.Client)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	if err := server.New(cfg.Server, client).Run(); err != nil {
		panic(err)
	}
}
