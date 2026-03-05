package loms

import (
	"context"

	pb "github.com/shizakira/cart/gen/grpc/loms_v1"
	"github.com/shizakira/cart/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Addr string `envconfig:"LOMS_GRPC_ADDR" default:"localhost:50051"`
}

type Service struct {
	client pb.LomsClient
}

func NewService(c Config) (*Service, error) {
	conn, err := grpc.NewClient(c.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Service{client: pb.NewLomsClient(conn)}, nil
}

func (c *Service) CreateOrder(ctx context.Context, userID int, items []domain.Item) (int64, error) {
	pbItems := make([]*pb.Item, 0, len(items))
	for _, item := range items {
		pbItems = append(pbItems, &pb.Item{
			Sku:   uint32(item.SkuID()),
			Count: uint32(item.Count()),
		})
	}

	resp, err := c.client.OrderCreate(ctx, &pb.OrderCreateRequest{
		User:  int64(userID),
		Items: pbItems,
	})
	if err != nil {
		return 0, err
	}

	return resp.GetOrderId(), nil
}

func (c *Service) StocksInfo(ctx context.Context, sku uint32) (uint64, error) {
	resp, err := c.client.StocksInfo(ctx, &pb.StocksInfoRequest{Sku: sku})
	if err != nil {
		return 0, err
	}

	return resp.GetCount(), nil
}
