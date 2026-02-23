package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Addr         string        `default:"localhost" envconfig:"HTTP_ADDR"`
	Port         string        `default:"8080" envconfig:"HTTP_PORT"`
	ReadTimeout  time.Duration `default:"20s" envconfig:"HTTP_READ_TIMEOUT"`
	WriteTimeout time.Duration `default:"20s" envconfig:"HTTP_WRITE_TIMEOUT"`
}

type Server struct {
	server *http.Server
}

func New(handler http.Handler, c Config) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		Addr:         net.JoinHostPort(c.Addr, c.Port),
	}

	s := &Server{
		server: httpServer,
	}

	go s.start()

	log.Info().Msgf("http server: started on: %s:%s", c.Addr, c.Port)

	return s
}

func (s *Server) start() {
	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("http server: ListenAndServe")
	}
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("http server: s.server.Shutdown")
	}

	log.Info().Msg("http server: closed")
}
