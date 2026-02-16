//go:build integration

package test

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/app"
	"github.com/shizakira/cart/internal/config"
	"github.com/shizakira/cart/pkg/httpserver"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var rootPath string = ""

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	*require.Assertions
	*http.Client
}

func (s *Suite) SetupSuite() {
	s.Assertions = s.Require()
	s.Client = &http.Client{}

	c := config.Config{
		App: config.App{
			Name:    "cart-service",
			Version: "test",
		},
		HTTP: httpserver.Config{
			Addr:         "localhost",
			Port:         "8083",
			ReadTimeout:  20 * time.Second,
			WriteTimeout: 20 * time.Second,
		},
	}

	addr := net.JoinHostPort(c.HTTP.Addr, c.HTTP.Port)
	rootPath = "http://" + addr

	log.Logger = zerolog.Nop()

	go func() {
		err := app.Run(context.Background(), c)
		s.NoError(err)
	}()

	time.Sleep(200 * time.Millisecond)
}
func (s *Suite) TearDownSuite() {}

func (s *Suite) SetupTest() {}

func (s *Suite) TearDownTest() {}
