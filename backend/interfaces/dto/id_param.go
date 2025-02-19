package dto

import "github.com/akira-saneyoshi/task-app/application"

type IDParam struct {
	id string
}

func NewIDParam(id string) *IDParam {
	return &IDParam{id}
}

func (i *IDParam) Value() string {
	return i.id
}

func (i *IDParam) Validate() error {
	if len(i.id) > 50 {
		return &application.ErrInputValidationFailed{Msg: "[ERROR] id must be 50 characters or less"}
	}
	return nil
}
