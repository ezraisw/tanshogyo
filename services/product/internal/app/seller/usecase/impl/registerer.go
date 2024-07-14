package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	"github.com/ezraisw/tanshogyo/pkg/common/util/idgen"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	sellererrors "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/errors"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

type SellerRegistererOptions struct {
	SellerFormValidator usecase.SellerFormValidator
	SellerRepository    repository.SellerRepository
	Nower               timehelper.Nower
	IDGen               idgen.IDGen
}

type SellerRegisterer struct {
	o SellerRegistererOptions
}

func NewSellerRegisterer(options SellerRegistererOptions) *SellerRegisterer {
	return &SellerRegisterer{
		o: options,
	}
}

func (s SellerRegisterer) Register(ctx context.Context, userId string, form usecase.SellerForm) error {
	if err := s.o.SellerFormValidator.Validate(ctx, form); err != nil {
		return err
	}

	if err := s.checkExistingUserId(ctx, userId); err != nil {
		return err
	}

	now := s.o.Nower.Now().UTC()

	seller := fromForm(form)
	seller.ID = s.o.IDGen.Generate()
	seller.UserID = userId
	seller.CreatedAt = now
	seller.UpdatedAt = now

	if _, err := s.o.SellerRepository.Create(ctx, seller); err != nil {
		return err
	}

	return nil
}

func (s SellerRegisterer) checkExistingUserId(ctx context.Context, userId string) error {
	exists, err := s.o.SellerRepository.Exists(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("UserID").Is(entity.OperatorEquals, userId)
	})
	if err != nil {
		return err
	}

	if exists {
		return sellererrors.ErrAlreadyHasSellerAccount
	}

	return nil
}
