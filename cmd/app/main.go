package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	bigfile2 "upload-big-file-to-elma/___backup/internal/app/bigfile"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/app.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := bigfile2.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	err = bigfile2.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
