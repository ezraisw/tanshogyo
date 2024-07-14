package app

import transactionweb "github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/adapter/web"

type HandlerRegistries struct {
	TransactionHTTP *transactionweb.TransactionHandlerRegistry
}
