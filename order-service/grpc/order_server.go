package grpc

import (
	"context"
	inventorypb "inventory-service/proto"
	"order-service/internal/entity"
	"order-service/internal/usecase"
	orderpb "order-service/proto"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	UC              usecase.OrderUseCase
	inventoryClient inventorypb.InventoryServiceClient
}

func NewOrderServer(uc usecase.OrderUseCase, inventoryClient inventorypb.InventoryServiceClient) *OrderServer {
	return &OrderServer{UC: uc, inventoryClient: inventoryClient}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	order := &entity.Order{
		ID:        primitive.NewObjectID(),
		UserID:    req.GetUserId(),
		Status:    req.GetStatus(),
		CreatedAt: time.Now(),
	}

	for _, item := range req.GetProducts() {
		productResp, err := s.inventoryClient.GetProductByID(ctx, &inventorypb.GetProductRequest{
			Id: item.GetProductId(),
		})
		if err != nil {
			return nil, err
		}

		order.Products = append(order.Products, entity.OrderItem{
			ProductID: item.GetProductId(),
			Quantity:  int(item.GetQuantity()),
			Price:     float64(productResp.GetPrice()), // convert from float32 to float64
		})
	}

	if err := s.UC.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	return toOrderResponse(order), nil
}

func (s *OrderServer) GetOrderByID(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.OrderResponse, error) {
	order, err := s.UC.GetOrderByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return toOrderResponse(order), nil
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.OrderResponse, error) {
	if err := s.UC.UpdateOrderStatus(ctx, req.GetId(), req.GetStatus()); err != nil {
		return nil, err
	}

	order, err := s.UC.GetOrderByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return toOrderResponse(order), nil
}

func (s *OrderServer) ListUserOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	orders, err := s.UC.GetOrdersByUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	var responses []*orderpb.OrderResponse
	for _, o := range orders {
		responses = append(responses, toOrderResponse(&o))
	}

	return &orderpb.ListOrdersResponse{Orders: responses}, nil
}

func toOrderResponse(order *entity.Order) *orderpb.OrderResponse {
	var items []*orderpb.OrderItem
	for _, p := range order.Products {
		items = append(items, &orderpb.OrderItem{
			ProductId: p.ProductID,
			Quantity:  int32(p.Quantity),
			Price:     p.Price, // must match double (float64)
		})
	}

	return &orderpb.OrderResponse{
		Id:        order.ID.Hex(),
		UserId:    order.UserID,
		Products:  items,
		Status:    order.Status,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
	}
}
