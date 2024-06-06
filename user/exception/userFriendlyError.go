package exception

import "fmt"

type UserFriendlyErrorInterface interface {
	error
}

type UserFriendlyError struct {
	err error
}

func (u UserFriendlyError) Error() string {
	return u.err.Error()
}

func InitError(text string, a ...any) UserFriendlyErrorInterface {
	return UserFriendlyError{err: fmt.Errorf(text, a...)}
}
