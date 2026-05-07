package main

import (
	"fmt"
	"github.com/ZoZo-182/gh-following-commits-feed/internal/config"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("startup failed: %v", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	_ = cfg
	return nil
}
