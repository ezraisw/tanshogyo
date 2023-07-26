package usecaseimpl_test

import (
	"time"

	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
)

var (
	dummyId  = "ffffffff-ffff-ffff-ffff-ffffffffffff"
	dummyNow = time.Now()

	dummyFields = []usecase.Field{
		{Name: "Username", Value: "exampleuser"},
	}

	dummyForm = usecase.UserForm{
		Username: "example1",
		Password: "anjaymaranjay",
		Email:    "example@example.org",
		Name:     "John Doe",
	}

	dummyUser = model.User{
		ID:        dummyId,
		Username:  dummyForm.Username,
		Password:  dummyForm.Password,
		Email:     dummyForm.Email,
		Name:      dummyForm.Name,
		CreatedAt: dummyNow.UTC(),
		UpdatedAt: dummyNow.UTC(),
	}

	dummyLoginForm = usecase.LoginForm{
		Username: dummyForm.Username,
		Password: dummyForm.Password,
	}

	dummyAuthentication = model.Authentication{
		Token:     dummyId,
		UserID:    dummyId,
		CreatedAt: dummyNow.UTC(),
		UpdatedAt: dummyNow.UTC(),
	}
)
