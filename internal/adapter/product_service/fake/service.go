package fake

import (
	"context"
	"math/rand"
	"time"

	"github.com/shizakira/cart/internal/model"
)

type Service struct {
	products map[int]model.Product
}

func New() *Service {
	return &Service{
		products: map[int]model.Product{
			2008:      {Name: "Гречка пропаренная, в пакетиках", Price: 16},
			2958025:   {Name: "Roxy Music. Stranded. Remastered Edition", Price: 1028},
			773297411: {Name: "Кроссовки Nike JORDAN", Price: 2202},
		},
	}
}

func randomDelay() {
	time.Sleep(time.Millisecond * time.Duration(50+rand.Intn(150)))
}

func (p *Service) IsProductExist(ctx context.Context, skuID int) (bool, error) {
	randomDelay()
	_, ok := p.products[skuID]
	return ok, nil
}

func (p *Service) GetProduct(ctx context.Context, skuID int) (model.Product, error) {
	randomDelay()
	prod, ok := p.products[skuID]
	if !ok {
		return model.Product{}, model.ErrProductNotFound
	}
	return prod, nil
}
