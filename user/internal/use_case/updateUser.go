package use_case

import (
	"app/user/internal/entity"
	"database/sql"
	"gorm.io/gorm"
)

type UpdateUserInput struct {
	Uid      *string
	Login    *string
	Password *string
}

type DB interface {
	Take(interface{}, ...interface{}) (tx *gorm.DB)
	Save(interface{}) *gorm.DB
	Begin(...*sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
}

type UpdateUserUseCase struct {
	db DB
}

func NewUpdateUserUseCase(db DB) UpdateUserUseCase {
	return UpdateUserUseCase{db: db}
}

func (u UpdateUserUseCase) Exec(updateUserInput UpdateUserInput) error {
	user := &entity.User{}

	u.db.Take(user, updateUserInput.Uid)

	if user.Uid.ID() == 0 {
		return nil
	}

	if updateUserInput.Login != nil {
		user.Login = *updateUserInput.Login
	}

	if updateUserInput.Password != nil {
		user.SetPassword(*updateUserInput.Password)
	}

	u.db.Begin()
	u.db.Save(user)
	u.db.Commit()

	return nil
}
