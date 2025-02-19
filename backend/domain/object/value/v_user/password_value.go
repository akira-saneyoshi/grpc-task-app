package value

import (
	"github.com/akira-saneyoshi/task-app/domain"
)

type Password struct {
	value string
}

func NewPassword(value string) *Password {
	return &Password{value}
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) Validate() error {
	if p.value == "" {
		return &domain.ErrValidationFailed{Msg: "[ERROR] password is empty"}
	}
	return nil
}
