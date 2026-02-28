package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/adapter/loms"
	"github.com/shizakira/cart/internal/adapter/postgres"
	"github.com/shizakira/cart/internal/adapter/product"
	pgpool "github.com/shizakira/cart/pkg/postgres"

	"github.com/shizakira/cart/internal/config"
	"github.com/shizakira/cart/internal/controller/http"
	"github.com/shizakira/cart/internal/usecase"
	"github.com/shizakira/cart/pkg/httpserver"
)

func Run(ctx context.Context, c config.Config) error {
	pgPool, err := pgpool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("pgpool.New: %w", err)
	}

	//storage := in_memory_storage.New()
	storage := postgres.New(pgPool)

	productSvc := product.NewService(c.ProductService)
	lomsSvc, err := loms.NewService(c.LomsService)
	if err != nil {
		return fmt.Errorf("loms_client_grpc.New: %w", err)
	}
	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	r := http.Router(uc)
	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("app: started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("app: got signal to stop")

	httpServer.Close()
	storage.Close()

	log.Info().Msg("app: stopped")

	return nil
}
