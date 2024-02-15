package repository

import (
	entity "github.com/ito30/cdp/entity"
)

type UserRepo interface {
	Get(id int) (*entity.User, error)
	Create(user entity.User) error
}
