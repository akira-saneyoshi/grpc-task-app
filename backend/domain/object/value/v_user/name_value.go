package value

import (
	"github.com/akira-saneyoshi/task-app/domain"
)

type Name struct {
	value string
}

func NewName(value string) *Name {
	return &Name{value}
}

func (e *Name) Value() string {
	return e.value
}

func (e *Name) Validate() error {
	if e.value == "" {
		return &domain.ErrValidationFailed{Msg: "[ERROR] user-name is empty"}
	}
	return nil
}
