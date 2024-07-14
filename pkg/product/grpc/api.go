package userauthgrpc

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/util/grpchelper"
	"github.com/ezraisw/tanshogyo/pkg/product"
	"github.com/ezraisw/tanshogyo/pkg/product/grpc/pb"
)

type GRPCProductAPI struct {
	client *ProductGRPCClient
}

func NewGRPCProductAPI(client *ProductGRPCClient) *GRPCProductAPI {
	return &GRPCProductAPI{client: client}
}

func (a GRPCProductAPI) Get(ctx context.Context, id string) (product.Product, error) {
	conn, err := a.client.Dial()
	if err != nil {
		return product.Product{}, err
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	p, err := client.Get(ctx, &pb.SingleProductRequest{
		Id: id,
	})
	if err != nil {
		return product.Product{}, grpchelper.HandleError(err)
	}

	return fromPb(p), nil
}

func (a GRPCProductAPI) Update(ctx context.Context, id string, form product.ProductForm) (product.Product, error) {
	conn, err := a.client.Dial()
	if err != nil {
		return product.Product{}, err
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	p, err := client.Update(ctx, &pb.UpdateProductRequest{
		Id:          id,
		SellerId:    form.SellerID,
		Name:        form.Name,
		Description: form.Description,
		Price:       form.Price,
		Quantity:    int64(form.Quantity),
	})
	if err != nil {
		return product.Product{}, grpchelper.HandleError(err)
	}

	return fromPb(p), nil
}

func fromPb(p *pb.Product) product.Product {
	return product.Product{
		ID:          p.Id,
		SellerID:    p.SellerId,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    int(p.Quantity),
		CreatedAt:   p.CreatedAt.AsTime(),
		UpdatedAt:   p.UpdatedAt.AsTime(),
	}
}
