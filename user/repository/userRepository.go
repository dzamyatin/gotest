package repository

import (
	"app/user/entity"
	"app/user/lib"
	"fmt"
)

type UserRepositoryInterface interface {
	Create(user entity.User) error
	GetById(uid string) (entity.User, error)
	GetAll() ([]entity.User, error)
}

type UserRepository struct {
}

func (u UserRepository) Create(user entity.User) error {
	_, err := lib.DB.Exec(`INSERT INTO user
    (uid, login, passwordHash)
VALUES
    (?, ?, ?)`,
		user.Uid,
		user.Login,
		user.PasswordHash,
	)

	return err
}

func (u UserRepository) GetById(uid string) (entity.User, error) {
	user := entity.User{}
	err := lib.DB.QueryRow(`SELECT uid, login, passwordHash
FROM user
WHERE uid = ?
LIMIT 1
`, uid).Scan(
		&user.Uid, &user.Login, &user.PasswordHash,
	)

	if err != nil {
		return user, fmt.Errorf("Can't get %w", err)
	}

	return user, nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	rows, err := lib.DB.Query(`SELECT uid, login, passwordHash
FROM user
`,
	)

	if err != nil {
		return []entity.User{}, err
	}

	var users []entity.User
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.Uid, &user.Login, &user.PasswordHash)
		users = append(users, user)
	}

	return users, nil
}

func InitUserRepository() UserRepositoryInterface {
	return UserRepository{}
}
