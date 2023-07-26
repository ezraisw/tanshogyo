package app

import transactionweb "github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/adapter/web"

type HandlerRegistries struct {
	TransactionHTTP *transactionweb.TransactionHandlerRegistry
}
