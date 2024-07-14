package app

import (
	productmodel "github.com/ezraisw/tanshogyo/services/product/internal/app/product/model"
	sellermodel "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/model"
)

// Models to be migrated as a table.
var models = []any{
	&productmodel.Product{},
	&sellermodel.Seller{},
}
