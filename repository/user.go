package repository

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var UserCache map[int]*User = map[int]*User{
	1: {
		ID:   1,
		Name: "John",
	},
	2: {
		ID:   2,
		Name: "Doe",
	},
}

var UserDB map[int]*User = map[int]*User{
	1: {
		ID:   1,
		Name: "John",
	},
	2: {
		ID:   2,
		Name: "Doe",
	},
}
