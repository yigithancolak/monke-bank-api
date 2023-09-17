package grpcapi

import (
	"context"
	"errors"
	"log"
	"time"

	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/pb"
	"github.com/yigithancolak/monke-bank-api/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	foundUser, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "error finding user: %s", err)
	}

	err = util.ComparePassword(req.GetPassword(), foundUser.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "wrong credentials: %s", err)
	}

	accessToken, _, err := server.tokenMaker.CreateToken(foundUser.ID, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating token: %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(foundUser.ID, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating token: %s", err)
	}

	_, err = server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       refreshPayload.UserID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Unix(refreshPayload.ExpiresAt, 0),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating session: %s", err)
	}

	resp := &pb.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	log.Printf("response: %v", resp)

	return resp, nil
}
