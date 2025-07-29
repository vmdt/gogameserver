package domain

type IUserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserById(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
