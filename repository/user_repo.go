package repository

import (
	"backend_github_trending/model"
	req2 "backend_github_trending/model/req"
	"context"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req2.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
}
