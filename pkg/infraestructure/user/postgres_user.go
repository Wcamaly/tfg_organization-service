package user

import (
	"tfg/organization-service/pkg/config/errors"
	"tfg/organization-service/pkg/domain/models"
	"tfg/organization-service/pkg/domain/user"
	"time"
)

var (
	ErrDomainConversion = errors.Define("userdto.toDomain_error")
)

type UserDto struct {
	ID        string    `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"lastLogin"`
}

type UserPostgres struct {
	ID        string    `db:"id"`
	UserName  string    `db:"username"`
	Password  string    `db:"password"`
	LastLogin time.Time `db:"lastLogin"`
}

func (p *UserPostgres) toDomain() (*user.User, error) {
	id, err := models.NewID(p.ID)
	if err != nil {
		return nil, err
	}
	return user.NewUser(id, p.UserName, p.Password, p.LastLogin), nil
}

func (u UserDto) toDomain() (*user.User, error) {
	id, err := models.NewID(u.ID)
	if err != nil {
		return nil, err
	}

	userName, err := models.NewRequiredString(u.UserName)
	if err != nil {
		return nil, errors.New(ErrDomainConversion, "missing username")
	}

	password, err := models.NewRequiredString(u.Password)
	if err != nil {
		return nil, errors.New(ErrDomainConversion, "missing password")
	}

	return user.NewUser(id, string(userName), string(password), nil), nil

}

func fromDomain(u *user.User) UserDto {
	return UserDto{
		ID:        string(u.Id()),
		UserName:  u.UserName(),
		Password:  u.Password(),
		LastLogin: u.LastLogin(),
	}
}
