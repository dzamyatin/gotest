package converter

import (
	api "app/user/api/user/proto"
	"app/user/internal/use_case"
)

func UserUpdateRequestToUpdateUserInput(request *api.UserUpdateRequest) use_case.UpdateUserInput {
	return use_case.UpdateUserInput{
		Uid:      &request.Uid,
		Login:    &request.Login,
		Password: &request.Password,
	}
}
