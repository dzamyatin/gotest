package use_case

type UpdateUserInput struct {
	Login string
}

type UpdateUserUseCase struct {
	//gorm *gorm.DB
}

func NewUpdateUserUseCase() UpdateUserUseCase {
	return UpdateUserUseCase{}
}

func (u UpdateUserUseCase) Exec(createUserInput CreateUserInput) error {

	return nil
}
