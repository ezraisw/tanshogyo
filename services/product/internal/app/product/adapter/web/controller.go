package web

import (
	"net/http"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/pkg/common/util/httphelper"
	"github.com/ezraisw/tanshogyo/pkg/userauth"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ProductControllerOptions struct {
	ProductLister        usecase.ProductLister
	ProductGetter        usecase.ProductGetter
	ProductCreator       usecase.ProductCreator
	ProductUpdater       usecase.ProductUpdater
	ProductDeleter       usecase.ProductDeleter
	ProductAuthedLister  usecase.ProductAuthedLister
	ProductAuthedCreator usecase.ProductAuthedCreator
	ProductAuthedUpdater usecase.ProductAuthedUpdater
	ProductAuthedDeleter usecase.ProductAuthedDeleter
}

type ProductController struct {
	o ProductControllerOptions
}

func NewProductController(options ProductControllerOptions) *ProductController {
	return &ProductController{
		o: options,
	}
}

func (c ProductController) ListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limit := helper.AssumeInt(r.URL.Query().Get("limit"))
	offset := helper.AssumeInt(r.URL.Query().Get("offset"))

	list, err := c.o.ProductLister.List(ctx, limit, offset)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, list)
}

func (c ProductController) GetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	product, err := c.o.ProductGetter.Get(ctx, chi.URLParam(r, "id"))
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, product)
}

func (c ProductController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var form usecase.ProductForm
	ctx := r.Context()

	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	product, err := c.o.ProductCreator.Create(ctx, form)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusCreated, product)
}

func (c ProductController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var form usecase.ProductForm
	ctx := r.Context()

	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	product, err := c.o.ProductUpdater.Update(ctx, chi.URLParam(r, "id"), form)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, product)
}

func (c ProductController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := c.o.ProductDeleter.Delete(ctx, chi.URLParam(r, "id")); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.RespondStatusCode(w, r, http.StatusOK)
}

func (c ProductController) AuthedListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	limit := helper.AssumeInt(r.URL.Query().Get("limit"))
	offset := helper.AssumeInt(r.URL.Query().Get("offset"))

	list, err := c.o.ProductAuthedLister.AuthedList(ctx, authUser.ID, limit, offset)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, list)
}

func (c ProductController) AuthedCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	var form usecase.AuthedProductForm

	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	product, err := c.o.ProductAuthedCreator.AuthedCreate(ctx, authUser.ID, form)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusCreated, product)
}

func (c ProductController) AuthedUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	var form usecase.AuthedProductForm

	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	product, err := c.o.ProductAuthedUpdater.AuthedUpdate(ctx, authUser.ID, chi.URLParam(r, "id"), form)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, product)
}

func (c ProductController) AuthedDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authUser := userauth.FromContext(ctx)
	if authUser == nil {
		// Should be handled by middleware.
		httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
		return
	}

	if err := c.o.ProductAuthedDeleter.AuthedDelete(ctx, authUser.ID, chi.URLParam(r, "id")); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.RespondStatusCode(w, r, http.StatusOK)
}
