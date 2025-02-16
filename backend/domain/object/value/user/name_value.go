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

func (n *Name) Value() string {
	return n.value
}

func (n *Name) Validate() error {
	if n.value == "" {
		return &domain.ErrValidationFailed{Msg: "name is empty"}
	}
	return nil
}
