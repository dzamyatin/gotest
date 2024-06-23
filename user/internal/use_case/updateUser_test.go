package use_case

import (
	"app/user/internal/entity"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"log"
	"testing"
)

type LoadNotEmptyUserFromDatabaseMatcher struct{}

func NewLoadNotEmptyUserFromDatabaseMatcher() LoadNotEmptyUserFromDatabaseMatcher {
	return LoadNotEmptyUserFromDatabaseMatcher{}
}

func (u LoadNotEmptyUserFromDatabaseMatcher) Matches(x interface{}) bool {
	user, ok := x.(*entity.User)

	if !ok {
		log.Fatal("expected user")
	}

	user.Uid = uuid.New()

	return true
}

func (u LoadNotEmptyUserFromDatabaseMatcher) String() string {
	return "Check user"
}

type AnyMatcher struct{}

func NewAnyMatcher() AnyMatcher {
	return AnyMatcher{}
}

func (u AnyMatcher) Matches(x interface{}) bool {
	return true
}

func (u AnyMatcher) String() string {
	return "Check any"
}

type SaveUserMatcher struct{}

func NewSaveUserMatcher() SaveUserMatcher {
	return SaveUserMatcher{}
}

func (u SaveUserMatcher) Matches(x interface{}) bool {
	user, ok := x.(*entity.User)

	if !ok {
		log.Fatal("expected user")
	}

	if user.Login != "login" {
		return false
	}

	if !user.CheckPassword("pwd") {
		return false
	}

	return true
}

func (u SaveUserMatcher) String() string {
	return "Check saved user"
}

func TestUpdateUserUseCase_Exec(t *testing.T) {
	mock := NewMockDB(gomock.NewController(t))

	id := uuid.New()
	idstr := id.String()
	login := "login"
	pwd := "pwd"
	input := UpdateUserInput{
		Uid:      &idstr,
		Login:    &login,
		Password: &pwd,
	}

	mock.EXPECT().Begin()
	mock.EXPECT().Commit()
	mock.EXPECT().Save(NewSaveUserMatcher())
	mock.EXPECT().Take(NewLoadNotEmptyUserFromDatabaseMatcher(), NewAnyMatcher())

	useCase := UpdateUserUseCase{db: mock}

	err := useCase.Exec(input)

	if err != nil {
		t.Errorf("Error while executing use case: %v", err)
	}

}
