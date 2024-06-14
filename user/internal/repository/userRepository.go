package repository

import (
	"app/user/internal/entity"
	"database/sql"
	"fmt"
)

type UserRepositoryInterface interface {
	Create(user entity.User) error
	GetById(uid string) (entity.User, error)
	GetAll() ([]entity.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func (u UserRepository) Create(user entity.User) error {
	_, err := u.db.Exec(`INSERT INTO users
    (uid, login, password_hash)
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
	err := u.db.QueryRow(`SELECT uid, login, password_hash
FROM users
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
	rows, err := u.db.Query(`SELECT uid, login, password_hash
FROM users
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

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}
