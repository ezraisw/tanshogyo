package usecaseimpl_test

import (
	"time"

	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/model"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	sellerusecase "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

var (
	dummyId       = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	dummySellerId = "ffffffff-ffff-ffff-ffff-ffffffffffff"
	dummyNow      = time.Now()

	dummyUserId = "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"

	dummyLimit  = 10
	dummyOffset = 0

	dummySeller = sellerusecase.Seller{
		ID:        dummySellerId,
		UserID:    dummyUserId,
		ShopName:  "Lol Shop",
		CreatedAt: dummyNow.UTC(),
		UpdatedAt: dummyNow.UTC(),
	}

	dummyForm = usecase.ProductForm{
		SellerID:    "ffffffff-ffff-ffff-ffff-ffffffffffff",
		Name:        "Testing",
		Description: "Testing description",
		Price:       69000,
		Quantity:    50,
	}

	dummyAuthedForm = usecase.AuthedProductForm{
		Name:        dummyForm.Name,
		Description: dummyForm.Description,
		Price:       dummyForm.Price,
		Quantity:    dummyForm.Quantity,
	}

	dummyProduct = model.Product{
		ID:          dummyId,
		SellerID:    dummyForm.SellerID,
		Name:        dummyForm.Name,
		Description: dummyForm.Description,
		Price:       dummyForm.Price,
		Quantity:    dummyForm.Quantity,
		CreatedAt:   dummyNow.UTC(),
		UpdatedAt:   dummyNow.UTC(),
	}

	dummyProductDto = usecase.Product{
		ID:          dummyId,
		SellerID:    dummyForm.SellerID,
		Name:        dummyForm.Name,
		Description: dummyForm.Description,
		Price:       dummyForm.Price,
		Quantity:    dummyForm.Quantity,
		CreatedAt:   dummyNow.UTC(),
		UpdatedAt:   dummyNow.UTC(),
	}

	dummyProducts = []*model.Product{
		{
			ID:          dummyId,
			SellerID:    dummyForm.SellerID,
			Name:        dummyForm.Name + " 1",
			Description: dummyForm.Description + " 1",
			Price:       dummyForm.Price + 50000,
			Quantity:    dummyForm.Quantity + 2,
			CreatedAt:   dummyNow.UTC(),
			UpdatedAt:   dummyNow.UTC(),
		},
		{
			ID:          dummyId,
			SellerID:    dummyForm.SellerID,
			Name:        dummyForm.Name + " 2",
			Description: dummyForm.Description + " 2",
			Price:       dummyForm.Price + 10000,
			Quantity:    dummyForm.Quantity + 1,
			CreatedAt:   dummyNow.UTC(),
			UpdatedAt:   dummyNow.UTC(),
		},
	}
)
