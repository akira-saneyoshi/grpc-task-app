package value

import (
	"github.com/akira-saneyoshi/task-app/domain"
)

type Title struct {
	value string
}

func NewTitle(value string) *Title {
	return &Title{value}
}

func (t *Title) Value() string {
	return t.value
}

func (t *Title) Validate() error {
	if t.value == "" {
		return &domain.ErrValidationFailed{Msg: "title is empty"}
	}
	return nil
}
