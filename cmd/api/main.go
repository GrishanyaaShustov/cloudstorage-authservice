// Package main is the entry point for the auth service API.
package main

import (
	"os"

	"github.com/GrishanyaaShustov/cloudstorage-authservice/internal/config"
	"github.com/GrishanyaaShustov/cloudstorage-authservice/pkg/logger"
)

func main() {
	cfg := config.MustLoad(os.Getenv("CONFIG_PATH"))
	log := logger.SetupLogger(cfg.Env)

	log.Info("Hello world")
}
