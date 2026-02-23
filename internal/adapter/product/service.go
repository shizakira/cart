package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/shizakira/cart/internal/model"
)

var (
	ErrUnexpectedStatus = errors.New("unexpected status")
)

type Config struct {
	BaseURL string        `envconfig:"PRODUCT_SERVICE_BASE_URL" required:"true"`
	Token   string        `envconfig:"PRODUCT_SERVICE_TOKEN" required:"true"`
	Timeout time.Duration `default:"3s" envconfig:"PRODUCT_SERVICE_TIMEOUT"`
}

type Service struct {
	baseURL string
	token   string
	client  *http.Client
}

func NewService(c Config) *Service {
	client := &http.Client{
		Timeout: c.Timeout,
		Transport: &RetryTransport{
			Base:       http.DefaultTransport,
			MaxRetries: 3,
		},
	}

	return &Service{
		baseURL: c.BaseURL,
		token:   c.Token,
		client:  client,
	}
}

func (s *Service) GetProduct(ctx context.Context, skuID int) (model.Product, error) {
	resp, err := s.doGetProductReq(ctx, skuID)
	if err != nil {
		return model.Product{}, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var r model.Product
		if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return model.Product{}, fmt.Errorf("json.NewDecoder: %w", err)
		}
		return r, nil

	case http.StatusNotFound:
		return model.Product{}, model.ErrProductNotFound

	default:
		return model.Product{}, fmt.Errorf("%d: %w", resp.StatusCode, ErrUnexpectedStatus)
	}
}

func (s *Service) doGetProductReq(ctx context.Context, skuID int) (*http.Response, error) {
	type getProductRequest struct {
		Token string `json:"token"`
		SKU   int    `json:"sku"`
	}

	reqBody, err := json.Marshal(getProductRequest{
		Token: s.token,
		SKU:   skuID,
	})
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/get_product", s.baseURL),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return s.client.Do(req)
}
