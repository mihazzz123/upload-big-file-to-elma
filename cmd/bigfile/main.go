package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/app/bigfile"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/bigfile.toml", "path to config file")
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
