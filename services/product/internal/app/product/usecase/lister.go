package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=lister.go -destination=mock/lister_mock.go -package=usecasemock

type ProductLister interface {
	List(ctx context.Context, limit, offset int) (ProductList, error)
}
