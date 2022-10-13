package model

import (
	"pikachu/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User ...
type User struct {
	UID      string   `json:"uid" gorm:"primaryKey"`
	Email    string   `json:"email"`
	Password Password `json:"password,omitempty"`
	Nick     *string  `json:"nick,omitempty"`
}

// ValidateNewUser ...
func (u *User) ValidateNewUser() bool {
	return u.Email != "" && !u.Password.IsEmpty() && u.Nick != nil
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
	u.Password = Password(bytes)
	return nil
}

// UpdateUser ...
func (u *User) UpdateUser(user *User) *User {
	u.Email = user.Email
	u.Nick = user.Nick
	return u
}

// AfterFind ...
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context
	login, _ := ctx.Value(util.LoginKey).(bool)
	if !login {
		u.Password = ""
	}
	return
}
