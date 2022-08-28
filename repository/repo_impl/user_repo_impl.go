package repo_impl

import (
	"backend_github_trending/banana"
	"backend_github_trending/db"
	"backend_github_trending/log"
	"backend_github_trending/model"
	req2 "backend_github_trending/model/req"
	"backend_github_trending/repository"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		Insert into users(user_id,email,password,role,full_name,created_at,updated_at)
		values(:user_id,:email,:password,:role,:full_name,:created_at,:updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}
	return user, nil
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{sql: sql}
}

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req2.ReqSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "Select * from users where email=$1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
