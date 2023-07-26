package app

import usermodel "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"

// Models to be migrated as a table.
var models = []any{
	&usermodel.Authentication{},
	&usermodel.User{},
}
