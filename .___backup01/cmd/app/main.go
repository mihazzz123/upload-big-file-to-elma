package main

import (
	"log"
	"upload-big-file-to-elma/config"
	"upload-big-file-to-elma/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
