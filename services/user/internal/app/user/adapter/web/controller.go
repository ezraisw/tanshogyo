package web

import (
	"net/http"

	"github.com/ezraisw/tanshogyo/pkg/common/util/httphelper"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
	"github.com/go-chi/render"
)

type UserControllerOptions struct {
	UserAuthenticator usecase.UserAuthenticator
	UserLoginer       usecase.UserLoginer
	UserRegisterer    usecase.UserRegisterer
}

type UserController struct {
	o UserControllerOptions
}

func NewUserController(options UserControllerOptions) *UserController {
	return &UserController{
		o: options,
	}
}

func (c UserController) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := r.Header.Get("X-Auth-Token")

	user, err := c.o.UserAuthenticator.Authenticate(ctx, token)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, user)
}

func (c UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var form usecase.LoginForm

	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	result, err := c.o.UserLoginer.Login(ctx, form)
	if err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.Respond(w, r, http.StatusOK, result)
}

func (c UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var form usecase.UserForm

	if err := render.DecodeJSON(r.Body, &form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	if err := c.o.UserRegisterer.Register(ctx, form); err != nil {
		httphelper.RespondError(w, r, err)
		return
	}

	httphelper.RespondStatusCode(w, r, http.StatusOK)
}
