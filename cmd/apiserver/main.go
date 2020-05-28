package main

import (
	"flag"
	"go-http-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
)

// TODO: need to set ABSOLUTE path to the .toml file
func init() {
	flag.StringVar(&confPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	conf := apiserver.NewConfig()
	_, err := toml.DecodeFile(confPath, conf)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(conf)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
