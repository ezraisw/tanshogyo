package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=uniquechecker.go -destination=mock/uniquechecker_mock.go -package=usecasemock

type UserUniqueChecker interface {
	CheckUnique(ctx context.Context, excludedId string, fields []Field) (bool, error)
}
