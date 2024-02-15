package postgres

import (
	"github.com/ito30/cdp/entity"
)

type userRepoDB struct {
}

func NewUserRepo() *userRepoDB {
	return &userRepoDB{}
}

func (r *userRepoDB) Get(id int) (*entity.User, error) {
	return entity.UserDB[id], nil
}

func (r *userRepoDB) Create(user entity.User) error {
	entity.UserDB[user.ID] = &user

	return nil
}
