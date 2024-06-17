package use_case

import (
	"app/user/internal/entity"
	"gorm.io/gorm"
)

type UpdateUserInput struct {
	Uid      *string
	Login    *string
	Password *string
}

type UpdateUserUseCase struct {
	gorm *gorm.DB
}

func NewUpdateUserUseCase(gorm *gorm.DB) UpdateUserUseCase {
	return UpdateUserUseCase{gorm: gorm}
}

func (u UpdateUserUseCase) Exec(updateUserInput UpdateUserInput) error {
	user := &entity.User{}
	u.gorm.Take(user, updateUserInput.Uid)

	if len(user.Uid) == 0 {
		return nil
	}

	err := u.gorm.Transaction(
		func(tx *gorm.DB) error {
			if updateUserInput.Login != nil {
				user.Login = *updateUserInput.Login
			}

			if updateUserInput.Password != nil {
				user.SetPassword(*updateUserInput.Password)
			}

			return nil
		},
	)

	return err
}
