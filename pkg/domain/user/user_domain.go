package user

import (
	"context"
	"tfg/organization-service/pkg/config/errors"
	"tfg/organization-service/pkg/domain/models"
	"time"
)

var (
	ErrorUserInternal = errors.Define("user.internal_error")
	ErrorUserExist    = errors.Define("user.user_exist")
)

type Repository interface {
	Count(ctx context.Context, filter *UserFilter) (int, error)
	Find(ctx context.Context, filter *UserFilter) ([]*User, error)
	FindById(ctx context.Context, id models.ID) (*User, error)
	Create(ctx context.Context, c *User) error
	Update(ctx context.Context, c *User) error
}

type User struct {
	ID        models.ID
	username  string
	password  string
	lastLogin time.Time
}

func NewUser(id models.ID, username string, password string, lastLogin time.Time) *User {
	return &User{
		ID: id, username: username, password: password, lastLogin: lastLogin,
	}
}

func (u *User) Id() models.ID {
	return u.ID
}

func (u *User) UserName() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) LastLogin() time.Time {
	return u.lastLogin
}
