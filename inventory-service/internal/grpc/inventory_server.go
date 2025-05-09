package grpc

import (
	"context"
	"inventory-service/internal/entity"
	"inventory-service/internal/usecase"
	pb "inventory-service/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
	UC usecase.ProductUseCase
}

func NewInventoryServer(uc usecase.ProductUseCase) *InventoryServer {
	return &InventoryServer{UC: uc}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	product := &entity.Product{
		ID:       primitive.NewObjectID(),
		Name:     req.GetName(),
		Category: req.GetCategory(),
		Stock:    int(req.GetStock()),
		Price:    req.GetPrice(),
	}

	err := s.UC.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Id:       product.ID.Hex(),
		Name:     product.Name,
		Category: product.Category,
		Stock:    int32(product.Stock),
		Price:    product.Price,
	}, nil
}

func (s *InventoryServer) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	product, err := s.UC.GetProductByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Id:       product.ID.Hex(),
		Name:     product.Name,
		Category: product.Category,
		Stock:    int32(product.Stock),
		Price:    product.Price,
	}, nil
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	product := &entity.Product{
		Name:     req.GetName(),
		Category: req.GetCategory(),
		Stock:    int(req.GetStock()),
		Price:    req.GetPrice(),
	}

	err = s.UC.UpdateProduct(ctx, objID, product)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Id:       req.GetId(),
		Name:     product.Name,
		Category: product.Category,
		Stock:    int32(product.Stock),
		Price:    product.Price,
	}, nil
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Empty, error) {
	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	err = s.UC.DeleteProduct(ctx, objID)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, _ *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	filter := map[string]interface{}{}
	limit := int64(100)
	offset := int64(0)

	products, err := s.UC.GetAllProducts(ctx, filter, limit, offset)
	if err != nil {
		return nil, err
	}

	var protoProducts []*pb.ProductResponse
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.ProductResponse{
			Id:       p.ID.Hex(),
			Name:     p.Name,
			Category: p.Category,
			Stock:    int32(p.Stock),
			Price:    p.Price,
		})
	}

	return &pb.ListProductsResponse{Products: protoProducts}, nil
}
