package grpcapi

import (
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
