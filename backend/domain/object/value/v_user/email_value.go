package value

import (
	"github.com/akira-saneyoshi/task-app/domain"
)

type Email struct {
	value string
}

func NewEmail(value string) *Email {
	return &Email{value}
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Validate() error {
	if e.value == "" {
		return &domain.ErrValidationFailed{Msg: "[ERROR] email is empty"}
	}
	return nil
}
