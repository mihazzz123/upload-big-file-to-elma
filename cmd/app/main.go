package main

import (
	"log"
	"upload-big-file-to-elma/config"
)

func main() {
	// Configuration
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(config)
}
