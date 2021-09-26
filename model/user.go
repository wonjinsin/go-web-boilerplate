package model

import "fmt"

// User ...
type User struct {
	UID   string `json:"uid" gorm:"primaryKey"`
	Email string `json:"email"`
	Nick  string `json:"nick"`
}

// ValidateNewUser ...
func (u *User) ValidateNewUser() bool {
	return u.Email != "" && u.Nick != ""
}

// ValidateUpdateUser ...
func (u *User) ValidateUpdateUser() bool {
	if u.UID != "" {
		return false
	}
	return true
}

// UpdateUser ...
func (u *User) UpdateUser(user *User) *User {
	fmt.Println(u)

	u.Email = user.Email
	u.Nick = user.Nick
	return u
}
