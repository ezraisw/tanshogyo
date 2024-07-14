package usecaseimpl

import (
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/model"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

func toDto(seller *model.Seller) usecase.Seller {
	return usecase.Seller{
		ID:        seller.ID,
		UserID:    seller.UserID,
		ShopName:  seller.ShopName,
		CreatedAt: seller.CreatedAt,
		UpdatedAt: seller.UpdatedAt,
	}
}

func fromForm(form usecase.SellerForm) *model.Seller {
	return &model.Seller{
		ShopName: form.ShopName,
	}
}
