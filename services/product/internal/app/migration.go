package app

import (
	productmodel "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	sellermodel "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/model"
)

// Models to be migrated as a table.
var models = []any{
	&productmodel.Product{},
	&sellermodel.Seller{},
}
