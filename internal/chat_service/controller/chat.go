package controller

import (
	"context"

	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/chat_service"
)

func (c *Controller) Healthz(ctx context.Context, req *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{},nil
}

func (c *Controller)Readyz(ctx context.Context, req *pb.ReadyzRequest) (*pb.ReadyzResponse, error) {
	return &pb.ReadyzResponse{},nil
}

//func (c *Controller) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
//	// 1. validation (in middleware)
//
//	// 2. convert delivery models to domain models/DTO
//	orderInfo := newOrderFromPbCreateOrderRequest(req)
//
//	// 3. call usecase
//	newOrder, err := c.OrdersUsecase.CreateOrder(ctx, orderInfo)
//	if err != nil {
//		return nil, err // обработается на уровне middleware
//	}
//
//	// 4. convert domain models/DTO to delivery models
//	response := &pb.CreateOrderResponse{
//		OrderId: newOrder.ID.String(),
//	}
//
//	// 5. return result
//	return response, nil
//}
