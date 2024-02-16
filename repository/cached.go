package repository

import (
	"errors"
	"fmt"
)

type cachedUserRepo struct {
	cache      map[int]*User
	userRepoDB UserRepo
}

func NewCachedUserRepo(userRepoDB UserRepo) *cachedUserRepo {
	return &cachedUserRepo{
		cache:      UserCache,
		userRepoDB: userRepoDB,
	}
}

func (r *cachedUserRepo) Get(id int) (*User, error) {
	cachedUser, ok := r.cache[id]
	if ok {
		fmt.Println("retrieving user from cache")
		return cachedUser, nil
	}

	user, err := r.userRepoDB.Get(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("record not found")
	}

	err = r.Create(*user)
	if err != nil {
		fmt.Println("error creating user cache:", err)
	}

	return user, nil
}

func (r *cachedUserRepo) Create(user User) error {
	if user.ID == 0 {
		return errors.New("bad request")
	}

	r.cache[user.ID] = &user
	return nil
}
