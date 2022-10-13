package model

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Auth ...
type Auth struct {
	Email    string   `json:"email"`
	Password Password `json:"password"`
}

// CheckPassword ...
func (a *Auth) CheckPassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(a.Password))
	return err == nil
}

// Validate ...
func (a *Auth) Validate() bool {
	return a.Email != "" && a.Password != ""
}

func (a *Auth) String() string {
	return fmt.Sprintf("Email[%s] Password[%s]",
		a.Email,
		fmt.Sprintf("%c%s%c", a.Password[0], strings.Repeat("*", len(a.Password)-2), a.Password[len(a.Password)-1]),
	)
}
