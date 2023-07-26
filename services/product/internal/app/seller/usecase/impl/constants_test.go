package usecaseimpl_test

import (
	"time"

	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/model"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
)

var (
	dummyId     = "ffffffff-ffff-ffff-ffff-ffffffffffff"
	dummyUserId = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	dummyNow    = time.Now()

	dummyForm = usecase.SellerForm{
		ShopName: "Super Deals Shop",
	}

	dummyFields = []usecase.Field{
		{Name: "ID", Value: dummyId},
	}

	dummySeller = model.Seller{
		ID:        dummyId,
		UserID:    dummyUserId,
		ShopName:  dummyForm.ShopName,
		CreatedAt: dummyNow.UTC(),
		UpdatedAt: dummyNow.UTC(),
	}
)
