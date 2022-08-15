package entity

import "sync"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	Users []User
	mu    sync.Mutex
}

var UsersDB Users

func (u *Users) Add(user User) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Users = append(u.Users, user)
}
