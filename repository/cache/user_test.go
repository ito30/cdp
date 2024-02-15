package cache

import (
	"testing"

	"github.com/ito30/cdp/entity"
	"github.com/ito30/cdp/repository"
	"github.com/ito30/cdp/repository/postgres"

	"github.com/stretchr/testify/assert"
)

func TestUserRepoCache_Get(t *testing.T) {
	tests := map[string]struct {
		userID int
		mock   func(userRepoDB repository.UserRepo)
		want   *entity.User
	}{
		"got from cache": {
			userID: 1,
			mock:   func(userRepoDB repository.UserRepo) {},
			want:   entity.Users[1],
		},
		"got from DB": {
			userID: 3,
			mock: func(userRepoDB repository.UserRepo) {
				userRepoDB.Create(entity.User{
					ID:   3,
					Name: "Jon",
				})
			},
			want: &entity.User{
				ID:   3,
				Name: "Jon",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repoDB := postgres.NewUserRepo()
			repoCache := NewUserRepo(repoDB)

			tt.mock(repoDB)

			got, _ := repoCache.Get(tt.userID)
			assert.Equal(t, tt.want, got)
		})
	}
}
