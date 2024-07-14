package web

import (
	"net/http"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/util/httphelper"
	"github.com/ezraisw/tanshogyo/pkg/userauth"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	"github.com/go-chi/render"
)

type SellerControllerOptions struct {
	SellerGetter   usecase.SellerGetter
	SellerRegister usecase.SellerRegisterer
}

type SellerController struct {
	o SellerControllerOptions
}

func NewSellerController(options SellerControllerOptions) *SellerController {
	return &SellerController{
		o: options,
	}
}

func (c SellerController) GetByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	user, err := c.o.SellerGetter.GetByUserID(ctx, authUser.ID)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, user)
}

func (c SellerController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	var form usecase.SellerForm
	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	if err := c.o.SellerRegister.Register(ctx, authUser.ID, form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.RespondStatusCode(w, r, http.StatusOK)
}
