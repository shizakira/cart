//go:build integration

package test

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/app"
	"github.com/shizakira/cart/internal/config"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var rootPath string

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

	godotenv.Overload(".env.test")
	c, err := config.NewTest()
	s.NoError(err)

	s.ResetMigrations()

	addr := net.JoinHostPort(c.HTTP.Addr, c.HTTP.Port)
	rootPath = "http://" + addr

	log.Logger = zerolog.Nop()

	go func() {
		err = app.Run(context.Background(), c)
		s.NoError(err)
	}()

	time.Sleep(200 * time.Millisecond)
}
