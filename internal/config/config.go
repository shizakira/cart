package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/shizakira/cart/internal/adapter/loms"
	"github.com/shizakira/cart/internal/adapter/product"
	"github.com/shizakira/cart/pkg/httpserver"
	"github.com/shizakira/cart/pkg/logger"
	"github.com/shizakira/cart/pkg/postgres"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App            App
	HTTP           httpserver.Config
	Logger         logger.Config
	Postgres       postgres.Config
	ProductService product.Config
	LomsService    loms.Config
}

func New() (Config, error) {
	return load(".env")
}

func NewTest() (Config, error) {
	return load(".env.test")
}

func load(env string) (Config, error) {
	var config Config

	if err := godotenv.Load(env); err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	if err := envconfig.Process("", &config); err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
