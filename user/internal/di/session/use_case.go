package session

import (
	"app/user/internal/use_case"
)

func (s *Session) GetUpdateUserUseCase() use_case.UpdateUserUseCase {
	return use_case.NewUpdateUserUseCase(s.GormSession())
}
