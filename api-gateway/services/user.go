package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/api-gateway/pb"
)

// GetUser - returns a user.
func (s *Service) GetUser(uID string) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userID := int64(rand.Intn(1000))
	res, err := s.UserSC.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return res, nil
}
