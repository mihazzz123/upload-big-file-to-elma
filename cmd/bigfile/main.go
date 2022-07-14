package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"upload-big-file-to-elma/internal/app/bigfile"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/bigfilestore.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := bigfile.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	err = bigfile.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
