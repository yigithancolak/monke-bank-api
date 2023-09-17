package grpcapi

import (
	"context"

	"github.com/google/uuid"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/pb"
	"github.com/yigithancolak/monke-bank-api/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)

	}

	arg := db.CreateUserParams{
		ID:       uuid.New(),
		Email:    req.GetEmail(),
		Password: hashedPassword,
		FullName: req.GetFullName(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// errCode := db.ErrorCode(err)
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	response := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return response, nil
}
