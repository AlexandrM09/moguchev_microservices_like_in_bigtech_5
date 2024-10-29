package middleware_grpc

import (
	"context"
	"errors"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorsUnaryServerInterceptor - convert any arror to rpc error
func ErrorsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			if _, ok := status.FromError(err); ok {
				return resp, err
			}
			switch {
			case errors.Is(err,models.ErrNotFoundFriend):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err,models.ErrWriteFriendFailed),
			 	errors.Is(err,models.ErrGetMessageFailed):
				return nil, status.Error(codes.Unavailable, err.Error())
			default:
				return nil, status.Error(codes.Internal, err.Error())
			}

		}

		return
	}
}
