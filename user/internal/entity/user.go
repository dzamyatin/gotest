package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uid          uuid.UUID
	Login        string
	PasswordHash []byte
}

func InitUser(
	login string,
) User {
	return User{
		Uid:   uuid.New(),
		Login: login,
	}
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))

	return err == nil
}

func (u User) SetPassword(password string) {
	u.PasswordHash, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}