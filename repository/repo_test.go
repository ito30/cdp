package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCachedUserRepo_Get(t *testing.T) {
	tests := map[string]struct {
		userID int
		repo   func() UserRepo
		want   *User
	}{
		"non-cached user repo": {
			userID: 1,
			repo: func() UserRepo {
				return NewNonCachedUserRepo()
			},
			want: UserCache[1],
		},
		"cached user repo": {
			userID: 1,
			repo: func() UserRepo {
				userRepo := NewNonCachedUserRepo()
				return NewCachedUserRepo(userRepo)
			},
			want: UserCache[1],
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := tt.repo().Get(tt.userID)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserRepoCache_Get(t *testing.T) {
	tests := map[string]struct {
		userID int
		mock   func(userRepoDB UserRepo)
		want   *User
	}{
		"got from cache": {
			userID: 1,
			mock:   func(userRepoDB UserRepo) {},
			want:   UserCache[1],
		},
		"got from DB": {
			userID: 3,
			mock: func(userRepoDB UserRepo) {
				userRepoDB.Create(User{
					ID:   3,
					Name: "Jon",
				})
			},
			want: &User{
				ID:   3,
				Name: "Jon",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			userRepo := NewNonCachedUserRepo()
			cachedUserRepo := NewCachedUserRepo(userRepo)

			tt.mock(userRepo)

			got, _ := cachedUserRepo.Get(tt.userID)
			assert.Equal(t, tt.want, got)
		})
	}
}
