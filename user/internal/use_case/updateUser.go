package use_case

import "gorm.io/gorm"

type UpdateUserInput struct {
	Uid      int
	Login    string
	Password string
}

type UpdateUserUseCase struct {
	gorm *gorm.DB
}

func NewUpdateUserUseCase(gorm *gorm.DB) UpdateUserUseCase {
	return UpdateUserUseCase{gorm: gorm}
}

func (u UpdateUserUseCase) Exec(createUserInput CreateUserInput) error {

	return nil
}
