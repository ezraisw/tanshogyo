package app

import transactionmodel "github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/model"

// Models to migrated as a table.
var models = []any{
	&transactionmodel.Transaction{},
	&transactionmodel.TransactionDetail{},
}
