package grpc

import (
	context "context"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/grpchelper"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/adapter/grpc/pb"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductServiceOptions struct {
	ProductGetter  usecase.ProductGetter
	ProductUpdater usecase.ProductUpdater
}

type ProductService struct {
	pb.UnimplementedProductServiceServer

	o ProductServiceOptions
}

func NewProductService(options ProductServiceOptions) *ProductService {
	return &ProductService{
		o: options,
	}
}

func (s ProductService) Get(ctx context.Context, req *pb.SingleProductRequest) (*pb.Product, error) {
	dto, err := s.o.ProductGetter.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(grpchelper.GetCode(err), err.Error())
	}

	msg := &pb.Product{
		Id:          dto.ID,
		SellerId:    dto.SellerID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Quantity:    int64(dto.Quantity),
		CreatedAt:   timestamppb.New(dto.CreatedAt),
		UpdatedAt:   timestamppb.New(dto.UpdatedAt),
	}
	return msg, nil
}

func (s ProductService) Update(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	form := usecase.ProductForm{
		SellerID:    req.SellerId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    int(req.Quantity),
	}

	dto, err := s.o.ProductUpdater.Update(ctx, req.Id, form)
	if err != nil {
		return nil, status.Error(grpchelper.GetCode(err), err.Error())
	}

	msg := &pb.Product{
		Id:          dto.ID,
		SellerId:    dto.SellerID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Quantity:    int64(dto.Quantity),
		CreatedAt:   timestamppb.New(dto.CreatedAt),
		UpdatedAt:   timestamppb.New(dto.UpdatedAt),
	}
	return msg, nil
}
