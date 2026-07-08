package repo

import "errors"

type User struct {
	ID          int    `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Store(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) Store(user User) (*User, error) {
	if user.ID == 0 {
		user.ID = len(u.users) + 1
	}

	u.users = append(u.users, user)
	return &u.users[len(u.users)-1], nil
}

func (u *userRepo) Find(email, pass string) (*User, error) {
	for idx := range u.users {
		if u.users[idx].Email == email && u.users[idx].Password == pass {
			return &u.users[idx], nil
		}
	}

	return nil, errors.New("user not found")
}
