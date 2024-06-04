package error

type UserFriendlyErrorInterface interface {
	error
}

type UserFriendlyError struct {
	text string
}

func (u UserFriendlyError) Error() string {
	return u.text
}

func InitError(text string) UserFriendlyErrorInterface {
	return UserFriendlyError{text: text}
}
