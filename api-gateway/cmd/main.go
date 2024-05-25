package main

import (
	"gateway/gateway"
	"gateway/server"
	"gateway/util"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	go gateway.RunGatewayServer(config)
	server.RunGrpcServer(config)
}
