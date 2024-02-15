package cache

import (
	"errors"
	"fmt"

	"github.com/ito30/cdp/entity"
	repo "github.com/ito30/cdp/repository"
)

type userRepoCache struct {
	cache      map[int]*entity.User
	userRepoDB repo.UserRepo
}

func NewUserRepo(userRepoDB repo.UserRepo) *userRepoCache {
	return &userRepoCache{
		cache:      entity.Users,
		userRepoDB: userRepoDB,
	}
}

func (r *userRepoCache) Get(id int) (*entity.User, error) {
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

func (r *userRepoCache) Create(user entity.User) error {
	if user.ID == 0 {
		return errors.New("bad request")
	}

	r.cache[user.ID] = &user
	return nil
}
