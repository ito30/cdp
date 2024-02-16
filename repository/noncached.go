package repository

import (
	"fmt"
)

type userRepo struct {
}

func NewNonCachedUserRepo() *userRepo {
	return &userRepo{}
}

func (r *userRepo) Get(id int) (*User, error) {
	fmt.Println("retrieving user from db")
	return UserDB[id], nil
}

func (r *userRepo) Create(user User) error {
	UserDB[user.ID] = &user

	return nil
}
