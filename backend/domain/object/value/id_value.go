package value

import (
	"github.com/akira-saneyoshi/task-app/domain"
)

type ID struct {
	value string
}

func NewID(value string) *ID {
	return &ID{value}
}

func (i *ID) Value() string {
	return i.value
}

func (i *ID) Validate() error {
	if i.value == "" {
		return &domain.ErrValidationFailed{Msg: "[ERROR] id is empty"}
	}
	return nil
}

func (i *ID) Equal(value string) bool {
	return i.value == value
}
