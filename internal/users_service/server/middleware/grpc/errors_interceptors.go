package middleware_grpc

import (
	"context"
	"errors"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
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
			case errors.Is(err,models.ErrAddFailed),
			 	 errors.Is(err,models.ErrRemoveFailed),
				 errors.Is(err,models.ErrConfirmFailed),
				 errors.Is(err,models.ErrGetListFailed):
				return nil, status.Error(codes.Unavailable, err.Error())
			default:
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
