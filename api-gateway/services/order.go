package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/api-gateway/utils"
	pb "github.com/voltgizerz/public-grpc/order/gen"
)

// GetOrder - returns an order.
func (s *Service) GetOrder(ordID string) (*pb.GetOrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer utils.RecoverPanic()

	orderID := int64(rand.Intn(100000000))
	res, err := s.OrderGRPC.GetOrder(ctx, &pb.GetOrderRequest{OrderId: orderID})
	if err != nil {
		return nil, err
	}

	return res, nil
}
