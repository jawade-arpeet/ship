package main

import (
	"ship/internal/config"
	"ship/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	srv, err := server.New()
	if err != nil {
		panic(err)
	}

	defer srv.Shutdown()

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
