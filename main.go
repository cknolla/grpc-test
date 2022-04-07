package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"grpc-test/server"
	"grpc-test/util"
	"net"
	"os"
	"os/signal"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Panic().Err(err).Msg("error loading config")
	}

	log.Info().Msgf("listening on %d", config.ListenPort)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.ListenHost, config.ListenPort))
	if err != nil {
		log.Panic().Err(err).Msg("failed to listen")
	}
	grpcServer := server.New()
	go func() {
		log.Info().Msg("starting server and listening for requests")
		if err := grpcServer.Serve(listener); err != nil {
			log.Panic().Err(err).Msg("failed to serve")
		}
	}()
	// wait for ctrl+c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Info().Msg("stopping server")
	grpcServer.Stop()
	log.Info().Msg("closing listener")
	_ = listener.Close()
}
