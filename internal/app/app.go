package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/adapter/in_memory_storage"
	fakeproductservice "github.com/shizakira/cart/internal/adapter/product_service/fake"
	"github.com/shizakira/cart/internal/config"
	"github.com/shizakira/cart/internal/controller/http"
	"github.com/shizakira/cart/internal/usecase"
	"github.com/shizakira/cart/pkg/httpserver"
)

func Run(ctx context.Context, c config.Config) error {
	storage := in_memory_storage.New()
	productSvc := fakeproductservice.New()

	uc := usecase.NewCart(storage, productSvc)

	r := http.Router(uc)
	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("app: started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("app: got signal to stop")

	httpServer.Close()

	log.Info().Msg("app: stopped")

	return nil
}
