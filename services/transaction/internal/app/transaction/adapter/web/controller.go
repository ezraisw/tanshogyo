package web

import (
	"net/http"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/pkg/common/util/httphelper"
	"github.com/ezraisw/tanshogyo/pkg/userauth"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/usecase"
	"github.com/go-chi/render"
)

type TransactionControllerOptions struct {
	TransactionLister      usecase.TransactionLister
	TransactionCreator     usecase.TransactionCreator
	TransactionCartGetter  usecase.TransactionCartGetter
	TransactionCartUpdater usecase.TransactionCartUpdater
}

type TransactionController struct {
	o TransactionControllerOptions
}

func NewTransactionController(options TransactionControllerOptions) *TransactionController {
	return &TransactionController{
		o: options,
	}
}

func (c TransactionController) ListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	limit := helper.AssumeInt(r.URL.Query().Get("limit"))
	offset := helper.AssumeInt(r.URL.Query().Get("offset"))

	list, err := c.o.TransactionLister.List(ctx, authUser.ID, limit, offset)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, list)
}

func (c TransactionController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	if err := c.o.TransactionCreator.Create(ctx, authUser.ID); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.RespondStatusCode(w, r, http.StatusCreated)
}

func (c TransactionController) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	cartInfo, err := c.o.TransactionCartGetter.GetCart(ctx, authUser.ID)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, cartInfo)
}

func (c TransactionController) UpdateCartHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	var cart usecase.Cart

	if err := render.DecodeJSON(r.Body, &cart); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	err := c.o.TransactionCartUpdater.UpdateCart(ctx, authUser.ID, cart)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.RespondStatusCode(w, r, http.StatusOK)
}
