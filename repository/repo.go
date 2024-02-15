package repository

import (
	"context"

	userentity "github.com/ito30/cdp/entity/user"
)

type UserRepo interface {
	Get(ctx context.Context, id int) (*userentity.User, error)
}
