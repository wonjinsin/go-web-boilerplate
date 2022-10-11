package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	UID      string  `json:"uid" gorm:"primaryKey"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Nick     *string `json:"nick,omitEmpty"`
}

// ValidateNewUser ...
func (u *User) ValidateNewUser() bool {
	return u.Email != "" && u.Password != "" && u.Nick != nil
}

// ValidateUpdateUser ...
func (u *User) ValidateUpdateUser() bool {
	if u.UID != "" {
		return false
	}
	return true
}

// UpdateHashPassword ...
func (u *User) UpdateHashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword ...
func (u *User) CheckPassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	return err == nil
}

// UpdateUser ...
func (u *User) UpdateUser(user *User) *User {
	u.Email = user.Email
	u.Nick = user.Nick
	return u
}
