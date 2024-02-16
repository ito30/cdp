package repository

type UserRepo interface {
	Get(id int) (*User, error)
	Create(user User) error
}
