package api

import (
	"context"
	"gateway/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RegisterUser(cxt context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := pb.User{
		Username:          req.GetUsername(),
		FullName:          req.GetFullName(),
		Email:             req.GetEmail(),
		PasswordChangedAt: timestamppb.New(time.Now()),
		CreatedAt:         timestamppb.New(time.Date(2019, 10, 2, 12, 22, 22, 121, time.UTC)),
	}

	return &pb.RegisterUserResponse{
		User: &user,
	}, nil
}
