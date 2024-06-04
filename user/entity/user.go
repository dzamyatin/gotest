package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	uid          uuid.UUID
	login        string
	passwordHash []byte
}

func (u User) getId() string {
	return u.uid.String()
}

func (u User) getLogin() string {
	return u.login
}

func InitUser(
	login string,
) User {
	return User{
		uid:   uuid.New(),
		login: login,
	}
}

func (u User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.passwordHash, []byte(password))

	return err == nil
}

func (u User) setPassword(password string) {
	u.passwordHash, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
