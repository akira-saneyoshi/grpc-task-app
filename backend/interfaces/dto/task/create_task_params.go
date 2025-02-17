package dto

import (
	"time"

	"github.com/akira-saneyoshi/task-app/application"
	"github.com/akira-saneyoshi/task-app/interfaces/dto"
)

type CreateTaskParams struct {
	userID      dto.IDParam
	title       string
	description string
	status      Status
	due_date    time.Time
}

func NewCreateTaskParams(userID string, title string) *CreateTaskParams {
	return &CreateTaskParams{
		userID:      *dto.NewIDParam(userID),
		title:       title,
		description: description,
		status:      status,
		due_date:    due_date,
	}
}

func (f *CreateTaskParams) UserID() string {
	return f.userID.Value()
}

func (f *CreateTaskParams) Title() string {
	return f.title
}

func (f *CreateTaskParams) Validate() error {
	if err := f.userID.Validate(); err != nil {
		return err
	}
	if len([]rune(f.title)) > 100 {
		return &application.ErrInputValidationFailed{Msg: "[ERROR] title must be 100 characters or less"}
	}
	return nil
}
